package handler

import (
	"fmt"
	"http-callback/server/http/middleware"
	api "http-callback/svcutil/api"
	"log"
	"net/http"
	"os"
)

type CallbackHandler struct {
	*Handler
}

func (h CallbackHandler) CreateDeployment(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	out, err := h.UC.Bash.ExecuteBash(ctx, os.Getenv("SCRIPT_PATH"))
	if err != nil {
		log.Println(err)
		api.JSONResponse(w, http.StatusBadRequest, http.StatusBadRequest, "Failed to execute deployment script.", nil, nil)
		return
	}

	fmt.Println(out)

	res := ctx.Value(middleware.ProcessKey).(string)

	api.JSONResponse(w, http.StatusOK, http.StatusOK, "Successfully execute deployment script.", res, nil)
}
