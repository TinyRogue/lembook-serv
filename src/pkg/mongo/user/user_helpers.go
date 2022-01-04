package user

import (
	"context"
	service "github.com/TinyRogue/lembook-serv/pkg/mongo"
	"github.com/TinyRogue/lembook-serv/pkg/user"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"unicode"
)

const minPasswordLen = 10

func FindUserBy(ctx context.Context, by string, value interface{}) (*user.User, error) {
	var res user.User
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
