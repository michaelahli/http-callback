package middleware

import (
	"net/http"
	"os"
)

func VerifyBasicAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, pass, _ := r.BasicAuth()
		if user != os.Getenv("BASIC_USERNAME") || pass != os.Getenv("BASIC_PASSWORD") {
			http.Error(w, "Invalid token", http.StatusBadRequest)
			return
		}

		next.ServeHTTP(w, r)
	})
}
