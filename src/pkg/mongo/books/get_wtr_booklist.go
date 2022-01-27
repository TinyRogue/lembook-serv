package books

import (
	"context"
	"github.com/TinyRogue/lembook-serv/cmd/gql/graph/generated/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (s *Service) GetWTRBooks(ctx context.Context, userID *string, page int64) (*model.UsersBooks, error) {
	_, _, wtrBooks, err := getUserBookLists(ctx, s, userID)
	if err != nil {
		return nil, err
	}

	var booksUIDs []string
	for _, bUID := range wtrBooks {
		booksUIDs = append(booksUIDs, *bUID)
	}

	if len(booksUIDs) == 0 {
		return &model.UsersBooks{}, nil
	}

	filter := bson.M{"uid": bson.M{"$in": booksUIDs}}
	var maxBooks int64 = 30
	skipBooks := maxBooks * page
	opts := options.FindOptions{
		Limit: &maxBooks,
		Skip:  &skipBooks,
	}

	cursor, err := s.BooksCollection.Find(ctx, filter, &opts)
	if err != nil {
		return nil, err
	}

	var bookRes = model.CategorizedBooks{
		Genre: "Te, które uwielbiasz",
		Books: nil,
	}

	if err := cursor.All(ctx, &bookRes.Books); err != nil {
		return nil, err
	}

	for _, book := range bookRes.Books {
		book.InList = WTR
	}

	var usersBooks model.UsersBooks
	usersBooks.Slices = append(usersBooks.Slices, &bookRes)

	return &usersBooks, nil
}
