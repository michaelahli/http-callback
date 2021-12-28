package middleware

import (
	"context"
	"errors"
	"http-callback/server/usecase"
	api "http-callback/svcutil/api"
	random "http-callback/svcutil/random"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-redis/redis"
)

func RegisterMiddleware(r *chi.Mux) {
	r.Use(recoverer)
	r.Use(notfound)
}

func recoverer(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rvr := recover(); rvr != nil {
				msg := "Internal Server Error!"
				api.JSONResponse(w, http.StatusInternalServerError, http.StatusInternalServerError, msg, []map[string]interface{}{}, errors.New("panic recovered"))
				return
			}
		}()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func notfound(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tctx := chi.NewRouteContext()
		rctx := chi.RouteContext(r.Context())

		if !rctx.Routes.Match(tctx, r.Method, r.URL.Path) {
			msg := "Request Not Found!"
			api.JSONResponse(w, http.StatusNotFound, http.StatusNotFound, msg, []map[string]interface{}{}, errors.New("error not found"))
			return
		}

		next.ServeHTTP(w, r)
	})
}

func RateLimiter(uc *usecase.UC) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var processId string

			if err := uc.GetFromRedis(CurrentProcess, &processId); err == nil {
				msg := "Please wait 1 minute before hitting the deployment api"
				api.JSONResponse(w, http.StatusUnprocessableEntity, http.StatusUnprocessableEntity, msg, []map[string]interface{}{}, errors.New("frequent api hit"))
				return
			} else if err != nil && err != redis.Nil {
				log.Println(err)
				msg := "redis error"
				api.JSONResponse(w, http.StatusInternalServerError, http.StatusInternalServerError, msg, []map[string]interface{}{}, err)
				return
			}

			processId = random.Generator(10)
			if err := uc.StoreToRedisExp(CurrentProcess, processId, "1m"); err != nil {
				log.Println(err)
				msg := "Failed to store process data to redis"
				api.JSONResponse(w, http.StatusBadRequest, http.StatusBadRequest, msg, []map[string]interface{}{}, err)
				return
			}

			ctx := context.WithValue(r.Context(), ProcessKey, processId)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
