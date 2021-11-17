package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/TinyRogue/lembook-serv/graph"
	"github.com/TinyRogue/lembook-serv/graph/generated"
	"github.com/TinyRogue/lembook-serv/internal/db"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func CorsMiddleware(next http.Handler, mode string) http.Handler {
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

func main() {
	_ = godotenv.Load()
	port := os.Getenv("PORT")
	var mode string
	if os.Args[1] == "--dev" {
		mode = "dev"
	} else {
		mode = "prod"
	}

	log.Printf("Server running in %s mode\n", mode)
	service.InitDb()
	defer service.Disconnect()
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", CorsMiddleware(srv, mode))
	log.Printf("GraphiQL http://localhost:%s/", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
