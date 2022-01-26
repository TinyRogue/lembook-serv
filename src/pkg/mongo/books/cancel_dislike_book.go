package books

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *Service) CancelDislikeTheBook(ctx context.Context, bookUID *string, userID *string) error {
	filter := bson.M{"uid": userID}
	update := bson.M{"$pull": bson.M{"dislikedBooks": bookUID}}
	_, err := s.UsersCollection.UpdateOne(ctx, filter, update)
	return err
}
