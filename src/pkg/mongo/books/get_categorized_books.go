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
	loved, disliked, wtr, err := getUserBookLists(ctx, s, userID)
	if err != nil {
		return nil, err
	}

	var usersBooks model.UsersBooks
	for _, g := range *genres {
		bookSlice, err := getBooksFrom(ctx, s, g, 0)
		if err != nil {
			return nil, err
		}
		for _, book := range bookSlice.Books {
			if strInSlice(book.UID, loved) {
				book.InList = LOVED
			} else if strInSlice(book.UID, disliked) {
				book.InList = DISLIKED
			} else if strInSlice(book.UID, wtr) {
				book.InList = WTR
			} else {
				book.InList = NONE
			}
		}
		usersBooks.Slices = append(usersBooks.Slices, bookSlice)
	}

	return &usersBooks, nil
}
