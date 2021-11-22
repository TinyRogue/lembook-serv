package middleware

import (
	"context"
	"encoding/json"
	"github.com/TinyRogue/lembook-serv/internal/models"
	"github.com/TinyRogue/lembook-serv/pkg/jwt"
	"log"
	"net/http"
)

type ContextKey string

const ContextUserKey ContextKey = "user"

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("auth")
		// Unauthenticated
		if header == "" {
			next.ServeHTTP(w, r)
			return
		}

		uid, err := jwt.ParseToken(header)
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}

		user, err := models.FindUserBy(r.Context(), "uid", uid)
		// Not exists
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}

		marshalledUser, err := json.Marshal(&user)
		if err != nil {
			log.Println("Couldn't marshal due to: ", err)
			next.ServeHTTP(w, r)
			return
		}

		ctx := context.WithValue(r.Context(), ContextUserKey, marshalledUser)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func FindUserByCtx(ctx context.Context) *models.User {
	raw, _ := ctx.Value(ContextUserKey).([]byte)
	var user models.User
	err := json.Unmarshal(raw, &user)
	if err != nil {
		log.Println("Couldn't unmarshall due to ", err)
		return nil
	}
	return &user
}
