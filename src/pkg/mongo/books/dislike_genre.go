package books

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *Service) DislikeGenre(ctx context.Context, genre *string, userID *string) error {
	filter := bson.M{"uid": userID}
	update := bson.M{"$pull": bson.M{"likedGenres": genre}}
	_, err := s.UsersCollection.UpdateOne(ctx, filter, update)
	return err
}
