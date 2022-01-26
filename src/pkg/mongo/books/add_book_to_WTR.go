package books

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *Service) AddBookToWTR(ctx context.Context, bookUID *string, userID *string) error {
	if err := s.CancelDislikeTheBook(ctx, bookUID, userID); err != nil {
		return err
	}
	if err := s.CancelLoveTheBook(ctx, bookUID, userID); err != nil {
		return err
	}
	_, err := s.UsersCollection.UpdateOne(ctx, bson.M{"uid": userID}, bson.D{{"$addToSet", bson.D{{"willingToRead", bookUID}}}})
	return err
}
