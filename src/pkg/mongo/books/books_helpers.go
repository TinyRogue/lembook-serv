package books

import (
	"context"
	"github.com/TinyRogue/lembook-serv/cmd/gql/graph/generated/model"
	service "github.com/TinyRogue/lembook-serv/pkg/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

func getAllGenres(ctx context.Context) (*model.Genres, error) {
	usersCollection := service.DB.Collection(service.GenresCollectionName)
	cursor, err := usersCollection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	var genres model.Genres
	if err := cursor.All(ctx, &genres.Genres); err != nil {
		return nil, err
	}
	return &genres, nil
}
