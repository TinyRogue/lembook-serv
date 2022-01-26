package books

import (
	"context"
	"github.com/TinyRogue/lembook-serv/cmd/gql/graph/generated/model"
)

func (s *Service) GetCategorizedBooks(ctx context.Context, userID *string) (*model.UsersBooks, error) {
	genres, err := getUserGenres(ctx, s, userID)
	if err != nil {
		return nil, err
	}

	var usersBooks model.UsersBooks

	for _, g := range *genres {
		bookSlice, err := getBooksFrom(ctx, s, g, 0)
		if err != nil {
			return nil, err
		}
		usersBooks.Slices = append(usersBooks.Slices, bookSlice)
	}

	return &usersBooks, nil
}
