package books

import (
	"context"
	"github.com/TinyRogue/lembook-serv/cmd/gql/graph/generated/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (s *Service) GetDislikedBooks(ctx context.Context, userID *string, page int64) (*model.UsersBooks, error) {
	_, disliked, _, err := getUserBookLists(ctx, s, userID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"uid.$in": disliked}
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

	var categorizedBooks = model.CategorizedBooks{
		Genre: "Te, za którymi nie płaczesz",
		Books: nil,
	}
	if err := cursor.All(ctx, &categorizedBooks.Books); err != nil {
		return nil, err
	}

	var usersBooks model.UsersBooks
	usersBooks.Slices = append(usersBooks.Slices, &categorizedBooks)

	return &usersBooks, nil
}