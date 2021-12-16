package middleware

import (
	"net/http"
	"os"
)

func Cors(next http.Handler, mode string) http.Handler {
	var origin string
	if mode == "dev" {
		origin = os.Getenv("DEV_ORIGIN")
	} else {
		origin = os.Getenv("PROD_ORIGIN")
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Request-Method", "POST, GET")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		next.ServeHTTP(w, r)
	})
}
