package books

import (
	"context"
	"github.com/TinyRogue/lembook-serv/cmd/gql/graph/generated/model"
	"go.mongodb.org/mongo-driver/bson"
)

func getAllGenres(ctx context.Context, s *Service) (*model.Genres, error) {
	cursor, err := s.GenresCollection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	var genres model.Genres
	if err := cursor.All(ctx, &genres.Genres); err != nil {
		return nil, err
	}
	return &genres, nil
}

func getUserGenres(ctx context.Context, s *Service, userUID *string) (*[]*string, error) {
	filter := bson.M{"uid": *userUID}
	var user model.User
	if err := s.UsersCollection.FindOne(ctx, filter).Decode(&user); err != nil {
		return nil, err
	}
	return &user.LikedGenres, nil
}
