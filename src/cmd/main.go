package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/TinyRogue/lembook-serv/cmd/gql/graph"
	"github.com/TinyRogue/lembook-serv/cmd/gql/graph/generated"
	"github.com/TinyRogue/lembook-serv/pkg/middleware"
	service "github.com/TinyRogue/lembook-serv/pkg/mongo"
	"github.com/TinyRogue/lembook-serv/pkg/mongo/books"
	"github.com/TinyRogue/lembook-serv/pkg/mongo/user"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	_ = godotenv.Load("../../../.env")
	port := os.Getenv("PORT")
	var mode string
	if len(os.Args) == 1 || os.Args[1] == "--dev" {
		mode = "dev"
	} else {
		mode = "prod"
	}

	log.Printf("Server running in %s mode\n", mode)
	service.InitDb()
	defer service.Disconnect()

	userService := user.Service{UsersCollection: service.DB.Collection(service.UsersCollectionName)}
	booksService := books.Service{
		UsersCollection:  service.DB.Collection(service.UsersCollectionName),
		BooksCollection:  service.DB.Collection(service.BooksCollectionName),
		GenresCollection: service.DB.Collection(service.GenresCollectionName),
		C: &http.Client{
			Timeout: 12 * time.Second,
		},
		PassKey:            os.Getenv("PASS_KEY"),
		PredictServiceAddr: os.Getenv("ML_SERVICE"),
	}

	if booksService.PassKey == "" || booksService.PredictServiceAddr == "" {
		log.Fatalln("Pass key, or predict service's address is not defined.")
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		UserService:  &userService,
		BooksService: &booksService,
	}}))
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", middleware.Cors(middleware.Auth(srv), mode))

	log.Printf("GraphiQL http://localhost:%s/", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
