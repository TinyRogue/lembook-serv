package books

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *Service) AddBookToWTR(ctx context.Context, bookUID *string, userID *string) error {
	_, err := s.UsersCollection.UpdateOne(ctx, bson.M{"uid": userID}, bson.D{{"$addToSet", bson.D{{"willingToRead", bookUID}}}})
	return err
}
