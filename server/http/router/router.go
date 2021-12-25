package http

import (
	"http-callback/server/handler"
	"http-callback/server/http/middleware"
	"http-callback/server/usecase"

	"github.com/go-chi/chi"
)

type RouteConfig struct {
	R  *chi.Mux
	UC *usecase.UC
}

func New(r *chi.Mux, uc *usecase.UC) *RouteConfig {
	return &RouteConfig{R: r, UC: uc}
}

func (rc *RouteConfig) RegisterMiddleware() { middleware.RegisterMiddleware(rc.R) }

func (rc *RouteConfig) RegisterRoutes() {
	hndlr := handler.New(rc.UC)
	rc.R.Route("/v1", func(r chi.Router) {
		r.Route("/deploy", func(r chi.Router) {
			r.Group(func(r chi.Router) {
				cbhndlr := handler.CallbackHandler{Handler: hndlr}
				r.Get("/", cbhndlr.Deploy)
			})
		})
	})

	rc.R.Route("/unimplemented", func(r chi.Router) {
		udhndlr := handler.UnimplementedHandler{Handler: hndlr}
		r.Get("/", udhndlr.ServeHTTP)
	})
}
