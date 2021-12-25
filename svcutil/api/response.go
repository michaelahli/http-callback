package json

import (
	"encoding/json"
	"net/http"
)

func JSONResponse(w http.ResponseWriter, httpCode int, statCode int, message string, payload interface{}, err error) {
	respPayload := map[string]interface{}{
		"code":    statCode,
		"message": message,
		"data":    payload,
		"error":   err,
	}

	response, _ := json.Marshal(respPayload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)
	w.Write(response)
}
