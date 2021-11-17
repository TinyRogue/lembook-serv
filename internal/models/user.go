package models

import (
	"context"
	"errors"
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
	InvalidPassword   = errors.New("invalid password")
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

	hashedPassword, err := hash.Generate(&req.GQLRegistration.Password)
	if err != nil {
		return err
	}

	UID, err := nano.Nanoid()
	newUser := User{
		UID:           UID,
		Username:      req.GQLRegistration.Username,
		Password:      hashedPassword,
		Token:         nil,
		TokenSelector: nil,
	}

	usersCollection := service.DB.Collection(service.UsersCollectionName)
	_, err = usersCollection.InsertOne(ctx, newUser)
	return nil
}

//func (user *User) Login() (*string, error) {
//	usersCollection := db.Client.Database("TheDB").Collection("Users")
//	var dbUser bson.M
//	if err := usersCollection.FindOne(context.TODO(), bson.M{"username": user.Username}).Decode(&dbUser); err != nil {
//		return nil, notExists
//	}
//	hash := fmt.Sprintf("%v", dbUser["password"])
//	if !checkPasswordHash(&user.Password, &hash) {
//		return nil, invalidPassword
//	}
//
//	token, err := jwt.GenerateToken(&user.Username)
//	if err != nil {
//		return nil, err
//	}
//
//	user.Token = *token
//	_, err = usersCollection.UpdateOne(context.TODO(), bson.M{"_id": dbUser["_id"]}, bson.M{"token": user.Token})
//	if err != nil {
//		return nil, err
//	}
//	return &user.Token, nil
//}
//
//TODO: write tests
//func generateUID(users *mongo.Collection, username *string) (*string, error) {
//	userCursor, err := users.Find(context.TODO(), bson.M{"username": username})
//	if err != nil {
//		return nil, err
//	}
//	defer func(userCursor *mongo.Cursor, ctx context.Context) {
//		_ = userCursor.Close(ctx)
//	}(userCursor, context.Background())
//
//	var uid string
//	for unique := false; !unique; {
//		unique = true
//		uid, err = nano.Generate("0123456789", 4)
//		if err != nil {
//			return nil, err
//		}
//
//		for userCursor.Next(context.TODO()) {
//			var currUsr bson.M
//			if err = userCursor.Decode(&currUsr); err != nil {
//				return nil, err
//			}
//			if currUsr["identifier"] == uid {
//				unique = false
//				break
//			}
//		}
//	}
//
//	return &uid, nil
//}

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
