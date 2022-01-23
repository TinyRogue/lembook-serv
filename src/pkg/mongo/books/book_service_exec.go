package books

import "go.mongodb.org/mongo-driver/mongo"

type Service struct {
	UsersCollection  *mongo.Collection
	BooksCollection  *mongo.Collection
	GenresCollection *mongo.Collection
}
