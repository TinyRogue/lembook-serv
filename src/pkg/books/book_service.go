package books

import (
	"context"
	"github.com/TinyRogue/lembook-serv/cmd/gql/graph/generated/model"
)

type Service interface {
	GetCategorizedBooks(ctx context.Context, userID *string) (*model.UsersBooks, error)
	GetGenres(ctx context.Context, userID *string) (*model.Genres, error)
	LikeGenre(ctx context.Context, genre *string, userID *string) error
	DislikeGenre(ctx context.Context, genre *string, userID *string) error
	LoveTheBook(ctx context.Context, bookUID *string, userID *string) error
	DislikeTheBook(ctx context.Context, bookUID *string, userID *string) error
	AddBookToWTR(ctx context.Context, bookUID *string, userID *string) error
	CancelLoveTheBook(ctx context.Context, bookUID *string, userID *string) error
	CancelDislikeTheBook(ctx context.Context, bookUID *string, userID *string) error
	CancelAddBookToWTR(ctx context.Context, bookUID *string, userID *string) error
}
