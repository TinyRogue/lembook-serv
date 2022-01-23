package books

import (
	"context"
	"github.com/TinyRogue/lembook-serv/cmd/gql/graph/generated/model"
)

type Service interface {
	FindBooks(ctx context.Context, userID *string) (model.UsersBooks, error)
	GetGenres(ctx context.Context, userID *string) (*model.Genres, error)
	LikeGenre(ctx context.Context, genre *string, userID *string) error
}
