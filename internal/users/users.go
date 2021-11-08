package users

import (
	"context"
	"errors"
	"fmt"
	"github.com/TinyRogue/lembook-serv/internal/db"
	"github.com/TinyRogue/lembook-serv/pkg/jwt"
	nano "github.com/matoous/go-nanoid/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var (
	notExists       = errors.New("user does not exist")
	invalidPassword = errors.New("invalid password")
)

type User struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	Identifier string `json:"identifier"`
	Token      string `json:"token"`
}

//TODO: write tests
func (user *User) Create() error {
	hashedPassword, err := hashPassword(&user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

	usersCollection := db.Client.Database("TheDB").Collection("Users")
	uid, err := generateUID(usersCollection, &user.Username)
	if err != nil {
		return err
	}

	user.Identifier = *uid
	_, err = usersCollection.InsertOne(context.TODO(), user)
	return err
}

func (user *User) Login() (*string, error) {
	usersCollection := db.Client.Database("TheDB").Collection("Users")
	var dbUser bson.M
	if err := usersCollection.FindOne(context.TODO(), bson.M{"username": user.Username}).Decode(&dbUser); err != nil {
		return nil, notExists
	}
	hash := fmt.Sprintf("%v", dbUser["password"])
	if !checkPasswordHash(&user.Password, &hash) {
		return nil, invalidPassword
	}

	token, err := jwt.GenerateToken(&user.Username)
	if err != nil {
		return nil, err
	}

	user.Token = *token
	_, err = usersCollection.UpdateOne(context.TODO(), bson.M{"_id": dbUser["_id"]}, bson.M{"token": user.Token})
	if err != nil {
		return nil, err
	}
	return &user.Token, nil
}

//TODO: write tests
func generateUID(users *mongo.Collection, username *string) (*string, error) {
	userCursor, err := users.Find(context.TODO(), bson.M{"username": username})
	if err != nil {
		return nil, err
	}
	defer func(userCursor *mongo.Cursor, ctx context.Context) {
		_ = userCursor.Close(ctx)
	}(userCursor, context.Background())

	var uid string
	for unique := false; !unique; {
		unique = true
		uid, err = nano.Generate("0123456789", 4)
		if err != nil {
			return nil, err
		}

		for userCursor.Next(context.TODO()) {
			var currUsr bson.M
			if err = userCursor.Decode(&currUsr); err != nil {
				return nil, err
			}
			if currUsr["identifier"] == uid {
				unique = false
				break
			}
		}
	}

	return &uid, nil
}

func hashPassword(password *string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(*password), 10)
	return string(bytes), err
}

func checkPasswordHash(password, hash *string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(*hash), []byte(*password))
	return err == nil
}
