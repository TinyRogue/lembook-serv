package models

import (
	"context"
	"errors"
	"github.com/TinyRogue/lembook-serv/graph/generated/model"
	service "github.com/TinyRogue/lembook-serv/internal/db"
	"github.com/TinyRogue/lembook-serv/pkg/hash"
	"github.com/TinyRogue/lembook-serv/pkg/jwt"
	nano "github.com/matoous/go-nanoid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"unicode"
)

const minPasswordLen = 10

var (
	UserAlreadyExists      = errors.New("user already exists")
	InvalidCredentials     = errors.New("invalid credentials")
	InvalidPasswordRequest = errors.New("password does not meet its requirements")
)

type Registration struct {
	GQLRegistration model.Registration `json:"gql_registration"`
}

type User struct {
	UID      string    `json:"uid"`
	Username string    `json:"username"`
	Password string    `json:"Password"`
	Token    []*string `json:"Token"`
}

func (req *Registration) Save(ctx context.Context) error {
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

	usersCollection := service.DB.Collection(service.UsersCollectionName)
	_, err = usersCollection.InsertOne(ctx, newUser)
	return err
}

func (u *User) Login(ctx context.Context) (*string, error) {
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

	token, err := dbUser.AssignNewToken(ctx)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (u *User) AssignNewToken(ctx context.Context) (*string, error) {
	token, err := jwt.GenerateToken(&u.UID)
	if err != nil {
		return nil, err
	}
	usersCollection := service.DB.Collection(service.UsersCollectionName)
	_, err = usersCollection.UpdateOne(ctx, bson.M{"uid": u.UID}, bson.D{{"$push", bson.D{{"token", token}}}})
	return token, err
}

func FindUserBy(ctx context.Context, by string, value interface{}) (*User, error) {
	var res User
	usersCollection := service.DB.Collection(service.UsersCollectionName)
	err := usersCollection.FindOne(ctx, bson.M{by: value}).Decode(&res)
	return &res, err
}

func IsUsernameTaken(ctx context.Context, username string) bool {
	_, err := FindUserBy(ctx, "username", username)
	return err != mongo.ErrNoDocuments
}

func IsPasswordValid(password string) bool {
	if len(password) < minPasswordLen {
		return false
	}

	var digit, lower, upper, sign bool

	for _, l := range password {
		switch {
		case unicode.IsDigit(l):
			digit = true
		case unicode.IsLower(l):
			lower = true
		case unicode.IsUpper(l):
			upper = true
		case unicode.IsPunct(l) || unicode.IsSymbol(l):
			sign = true
		}
	}
	return digit && lower && sign && upper
}
