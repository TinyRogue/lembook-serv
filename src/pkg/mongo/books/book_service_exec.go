package books

import (
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

type Service struct {
	UsersCollection    *mongo.Collection
	BooksCollection    *mongo.Collection
	GenresCollection   *mongo.Collection
	C                  *http.Client
	PredictServiceAddr string
	PassKey            string
}
