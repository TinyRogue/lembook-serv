package books

import (
	"context"
	"github.com/TinyRogue/lembook-serv/cmd/gql/graph/generated/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	NONE     = 0
	LOVED    = 1
	DISLIKED = 2
	WTR      = 3
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

func getUserBookLists(ctx context.Context, s *Service, userUID *string) (loved []*string, disliked []*string, WTR []*string, err error) {
	filter := bson.M{"uid": *userUID}
	var user model.User
	if err := s.UsersCollection.FindOne(ctx, filter).Decode(&user); err != nil {
		return nil, nil, nil, err
	}
	return user.LikedBooks, user.DislikedBooks, user.WillingToRead, nil
}

func getUserGenres(ctx context.Context, s *Service, userUID *string) (*[]*string, error) {
	filter := bson.M{"uid": *userUID}
	var user model.User
	if err := s.UsersCollection.FindOne(ctx, filter).Decode(&user); err != nil {
		return nil, err
	}
	return &user.LikedGenres, nil
}

func getBooksFrom(ctx context.Context, s *Service, genre *string, page int64) (*model.CategorizedBooks, error) {
	var maxBooks int64 = 30
	skipBooks := maxBooks * page
	filter := bson.M{"genres": *genre}
	opts := options.FindOptions{
		Limit: &maxBooks,
		Skip:  &skipBooks,
	}

	cursor, err := s.BooksCollection.Find(ctx, filter, &opts)
	if err != nil {
		return nil, err
	}

	var categorizedBooks = model.CategorizedBooks{
		Genre: *genre,
		Books: nil,
	}

	if err := cursor.All(ctx, &categorizedBooks.Books); err != nil {
		return nil, err
	}
	return &categorizedBooks, nil
}

func strInSlice(sought string, in []*string) bool {
	for _, val := range in {
		if *val == sought {
			return true
		}
	}
	return false
}
