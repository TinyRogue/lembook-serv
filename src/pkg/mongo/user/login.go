package user

import (
	"context"
	"github.com/TinyRogue/lembook-serv/pkg/hash"
	"github.com/TinyRogue/lembook-serv/pkg/jwt"
	"github.com/TinyRogue/lembook-serv/pkg/user"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *Service) Login(ctx context.Context, u *user.User) (*string, error) {
	dbUser, err := FindUserBy(ctx, "username", u.Username)
	if err == mongo.ErrNoDocuments {
		return nil, user.InvalidCredentials
	} else if err != nil {
		return nil, err
	}

	beautifiedPassword := dbUser.Password
	match, err := hash.Compare(u.Password, beautifiedPassword)
	if err != nil {
		return nil, err
	}
	if !match {
		return nil, user.InvalidCredentials
	}

	token, err := s.AssignNewToken(ctx, dbUser)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (s *Service) AssignNewToken(ctx context.Context, u *user.User) (*string, error) {
	token, err := jwt.GenerateToken(&u.UID)
	if err != nil {
		return nil, err
	}
	_, err = s.UsersCollection.UpdateOne(ctx, bson.M{"uid": u.UID}, bson.D{{"$push", bson.D{{"token", token}}}})
	return token, err
}
