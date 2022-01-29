package books

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/TinyRogue/lembook-serv/cmd/gql/graph/generated/model"
	"go.mongodb.org/mongo-driver/bson"
	"io"
	"log"
)

func (s *Service) GetPredictedBooks(ctx context.Context, userID *string, booksNum int64) (*model.UsersBooks, error) {
	query := fmt.Sprintf("%s/similar-books/%s?key=%s&books=%d", s.PredictServiceAddr, *userID, s.PassKey, booksNum)
	log.Println(query)

	resp, err := s.C.Get(query)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("Error occurred during closing response body: %v\n", err.Error())
		}
	}(resp.Body)

	var predictedUIDs []string
	if err := json.NewDecoder(resp.Body).Decode(&predictedUIDs); err != nil {
		return nil, err
	}

	filter := bson.M{"uid": bson.M{"$in": predictedUIDs}}
	cursor, err := s.BooksCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	var bookRes = model.CategorizedBooks{
		Genre: "Starannie dobrane dla Ciebie",
		Books: nil,
	}

	if err := cursor.All(ctx, &bookRes.Books); err != nil {
		return nil, err
	}

	loved, disliked, wtr, err := getUserBookLists(ctx, s, userID)
	if err != nil {
		return nil, err
	}
	for _, book := range bookRes.Books {
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

	var usersBooks model.UsersBooks
	usersBooks.Slices = append(usersBooks.Slices, &bookRes)
	return &usersBooks, nil
}
