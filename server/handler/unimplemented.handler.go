package handler

import (
	api "http-callback/svcutil/api"
	"net/http"
)

type UnimplementedHandler struct {
	*Handler
}

func (h *UnimplementedHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	api.JSONResponse(w, http.StatusOK, http.StatusOK, "Succesfully Processing Unimplemented Request", nil, nil)
}
