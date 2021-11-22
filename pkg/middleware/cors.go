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
		w.Header().Set("Access-Control-Request-Method", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")

		next.ServeHTTP(w, r)
	})
}
