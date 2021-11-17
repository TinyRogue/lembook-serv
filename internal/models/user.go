package models

import (
	"context"
	"errors"
	"fmt"
	"github.com/TinyRogue/lembook-serv/graph/generated/model"
	service "github.com/TinyRogue/lembook-serv/internal/db"
	"github.com/TinyRogue/lembook-serv/pkg/hash"
	nano "github.com/matoous/go-nanoid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"unicode"
)

const minPasswordLen = 10

var (
	UserAlreadyExists = errors.New("user already exists")
	UserDoesntExists  = errors.New("user does not exists in database")
	InvalidPassword   = errors.New("password does not meet its requirements")
)

type Registration struct {
	GQLRegistration model.Registration `json:"gql_registration"`
}

type User struct {
	UID           string  `json:"uid"`
	Username      string  `json:"username"`
	Password      string  `json:"Password"`
	Token         *string `json:"Token"`
	TokenSelector *string `json:"TokenSelector"`
}

func (req *Registration) Save(ctx context.Context) error {
	if IsUsernameTaken(&ctx, req.GQLRegistration.Username) {
		return UserAlreadyExists
	}

	hashedPassword, err := hash.BeautifyPassword(req.GQLRegistration.Password, nil)
	if err != nil {
		return err
	}

	UID, _ := nano.Nanoid()
	newUser := User{
		UID:           UID,
		Username:      req.GQLRegistration.Username,
		Password:      hashedPassword,
		Token:         nil,
		TokenSelector: nil,
	}

	usersCollection := service.DB.Collection(service.UsersCollectionName)
	_, err = usersCollection.InsertOne(ctx, newUser)
	return err
}

func (user *User) Login() (*string, error) {
	usersCollection := service.DB.Collection("Users")
	var dbUser bson.M
	if err := usersCollection.FindOne(context.TODO(), bson.M{"username": user.Username}).Decode(&dbUser); err != nil {
		return nil, UserDoesntExists
	}
	beautifiedPassword := fmt.Sprintf("%v", dbUser["password"])
	match, err := hash.Compare(user.Password, beautifiedPassword)
	if err != nil {
		return nil, err
	}
	if !match {
		return nil, InvalidPassword
	}

	//token, err := jwt.GenerateToken(&user.Username)
	//if err != nil {
	//	return nil, err
	//}

	//user.Token = *token
	//_, err = usersCollection.UpdateOne(context.TODO(), bson.M{"_id": dbUser["_id"]}, bson.M{"token": user.Token})
	//if err != nil {
	//	return nil, err
	//}
	//return &user.Token, nil
	return nil, nil
}

func FindUserBy(ctx *context.Context, by string, value interface{}) (*User, error) {
	var res User
	usersCollection := service.DB.Collection(service.UsersCollectionName)
	err := usersCollection.FindOne(*ctx, bson.M{by: value}).Decode(&res)
	return &res, err
}

func IsUsernameTaken(ctx *context.Context, username string) bool {
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
