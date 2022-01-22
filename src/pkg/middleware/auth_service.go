package middleware

import (
	"context"
	"encoding/json"
	"github.com/TinyRogue/lembook-serv/cmd/gql/graph/generated/model"
	"github.com/TinyRogue/lembook-serv/pkg/jwt"
	"github.com/TinyRogue/lembook-serv/pkg/mongo/user"
	"log"
	"net/http"
)

type ContextKey string

const ContextUserKey ContextKey = "user"
const ContextReqWriterKey ContextKey = "reqWriter"
const ContextJWT ContextKey = "jwt"

type JWTHandler struct {
	jwt string
}

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, _ := r.Cookie("auth")
		ctx := context.WithValue(r.Context(), ContextReqWriterKey, &w)
		r = r.WithContext(ctx)

		// Unauthenticated
		if cookie == nil {
			log.Println("User unauthenticated")
			next.ServeHTTP(w, r)
			return
		}

		ctx = context.WithValue(r.Context(), ContextJWT, &JWTHandler{jwt: cookie.Value})
		r = r.WithContext(ctx)

		uid, err := jwt.ParseToken(cookie.Value)
		if err != nil {
			log.Printf("Parsing error %s\n", err.Error())
			next.ServeHTTP(w, r)
			return
		}

		u, err := user.FindUserBy(r.Context(), "uid", uid)
		// Not exists
		if err != nil {
			log.Printf("User not found by uid --> %s\n", err.Error())
			next.ServeHTTP(w, r)
			return
		}

		marshalledUser, err := json.Marshal(u)
		if err != nil {
			log.Printf("User could not be marshalled --> %s\n", err.Error())
			next.ServeHTTP(w, r)
			return
		}

		userCtx := context.WithValue(r.Context(), ContextUserKey, &marshalledUser)
		r = r.WithContext(userCtx)
		next.ServeHTTP(w, r)
	})
}

func FindUserByCtx(ctx context.Context) *model.User {
	raw := ctx.Value(ContextUserKey)
	if raw == nil {
		log.Println("No user key")
		return nil
	}

	var u model.User
	err := json.Unmarshal(*raw.(*[]byte), &u)
	if err != nil {
		log.Println("Couldn't unmarshall user due to ", err)
		return nil
	}
	return &u
}

func GetResWriter(ctx context.Context) *http.ResponseWriter {
	resWriter, ok := ctx.Value(ContextReqWriterKey).(*http.ResponseWriter)
	if !ok {
		return nil
	}
	return resWriter
}
