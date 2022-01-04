package user

import (
	"context"
	"errors"
	"github.com/TinyRogue/lembook-serv/cmd/gql/graph/generated/model"
	"github.com/TinyRogue/lembook-serv/pkg/hash"
	"github.com/TinyRogue/lembook-serv/pkg/jwt"
	nano "github.com/matoous/go-nanoid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	UserAlreadyExists      = errors.New("user already exists")
	InvalidCredentials     = errors.New("invalid credentials")
	InvalidPasswordRequest = errors.New("password does not meet its requirements")
)

type Service struct {
	UsersCollection *mongo.Collection
}

type User struct {
	UID      string    `json:"uid"`
	Username string    `json:"username"`
	Password string    `json:"Password"`
	Token    []*string `json:"Token"`
}

type Registration struct {
	GQLRegistration model.Registration `json:"gql_registration"`
}

func (s *Service) Login(ctx context.Context, u *User) (*string, error) {
	dbUser, err := FindUserBy(ctx, "username", u.Username)
	if err == mongo.ErrNoDocuments {
		return nil, InvalidCredentials
	} else if err != nil {
		return nil, err
	}

	beautifiedPassword := dbUser.Password
	match, err := hash.Compare(u.Password, beautifiedPassword)
	if err != nil {
		return nil, err
	}
	if !match {
		return nil, InvalidCredentials
	}

	token, err := s.AssignNewToken(ctx, dbUser)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (s *Service) AssignNewToken(ctx context.Context, u *User) (*string, error) {
	token, err := jwt.GenerateToken(&u.UID)
	if err != nil {
		return nil, err
	}
	_, err = s.UsersCollection.UpdateOne(ctx, bson.M{"uid": u.UID}, bson.D{{"$push", bson.D{{"token", token}}}})
	return token, err
}

func (s *Service) Register(ctx context.Context, req *Registration) error {
	if IsUsernameTaken(ctx, req.GQLRegistration.Username) {
		return UserAlreadyExists
	}

	hashedPassword, err := hash.BeautifyPassword(req.GQLRegistration.Password, nil)
	if err != nil {
		return err
	}

	UID, _ := nano.Nanoid()
	newUser := User{
		UID:      UID,
		Username: req.GQLRegistration.Username,
		Password: hashedPassword,
		Token:    []*string{},
	}

	_, err = s.UsersCollection.InsertOne(ctx, newUser)
	return err
}
