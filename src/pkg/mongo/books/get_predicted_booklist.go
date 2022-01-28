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
	//http://localhost:8081/similar-books/ieqmxvyj-sYcbeHfafqSB?key=ITS_THE_PASS_KEY_BRUH_NOT_SO_SAFE$2$1&books=109
	//http://localhost:8081/similar-books/ieqmxvyj-sYcbeHfafqSB?key=ITS_THE_PASS_KEY_BRUH_NOT_SO_SAFE&books=20
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

	var usersBooks model.UsersBooks
	usersBooks.Slices = append(usersBooks.Slices, &bookRes)
	return &usersBooks, nil
}
