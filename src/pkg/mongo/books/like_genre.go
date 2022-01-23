package books

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *Service) LikeGenre(ctx context.Context, genre *string, userID *string) error {
	_, err := s.UsersCollection.UpdateOne(ctx, bson.M{"uid": userID}, bson.D{{"$push", bson.D{{"likedGenres", genre}}}})
	return err
}
