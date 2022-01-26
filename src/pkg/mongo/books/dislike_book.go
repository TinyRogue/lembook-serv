package books

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *Service) DislikeTheBook(ctx context.Context, bookUID *string, userID *string) error {
	if err := s.CancelAddBookToWTR(ctx, bookUID, userID); err != nil {
		return err
	}
	if err := s.CancelLoveTheBook(ctx, bookUID, userID); err != nil {
		return err
	}
	_, err := s.UsersCollection.UpdateOne(ctx, bson.M{"uid": userID}, bson.D{{"$addToSet", bson.D{{"dislikedBooks", bookUID}}}})
	return err
}
