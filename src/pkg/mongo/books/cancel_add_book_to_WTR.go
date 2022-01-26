package books

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *Service) CancelAddBookToWTR(ctx context.Context, bookUID *string, userID *string) error {
	filter := bson.M{"uid": userID}
	update := bson.M{"$pull": bson.M{"willingToRead": bookUID}}
	_, err := s.UsersCollection.UpdateOne(ctx, filter, update)
	return err
}
