package handler

import (
	api "http-callback/svcutil/api"
	"http-callback/svcutil/slice"
	"net/http"
	"os"
	"strings"
)

type CallbackHandler struct {
	*Handler
}

func (h CallbackHandler) Deploy(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	out, err := h.UC.Bash.ExecuteBash(ctx, os.Getenv("SCRIPT_PATH"))
	if err != nil {
		api.JSONResponse(w, http.StatusBadRequest, http.StatusBadRequest, "Failed to execute deployment script.", nil, nil)
		return
	}

	res := strings.Split(string(out), "\n")
	for i, s := range res {
		if len(s) == 0 {
			res = slice.RemoveElementFromStringArray(res, res[i])
		}
	}
	api.JSONResponse(w, http.StatusOK, http.StatusOK, "Successfully execute deployment script.", res, nil)
}
