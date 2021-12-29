package middleware

import (
	"context"
	"errors"
	"http-callback/server/usecase"
	api "http-callback/svcutil/api"
	random "http-callback/svcutil/random"
	"log"
	"net/http"

	"github.com/go-redis/redis"
)

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
