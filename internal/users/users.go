package users

import (
	"context"
	"github.com/TinyRogue/lembook-serv/internal/db"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username string `json:"name"`
	Password string `json:"password"`
}

func (user *User) Create() (*mongo.InsertOneResult, error) {
	hashedPassword, err := HashPassword(user.Password)
	user.Password = hashedPassword
	userID, err := db.Client.Database("TheDb").Collection("Users").InsertOne(context.TODO(), user)
	return userID, err
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
