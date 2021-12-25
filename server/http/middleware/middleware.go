package middleware

import (
	"errors"
	api "http-callback/svcutil/api"
	"net/http"

	"github.com/go-chi/chi"
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
