package user

import (
	"errors"
	"github.com/TinyRogue/lembook-serv/cmd/gql/graph/generated/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type Service struct {
	UsersCollection *mongo.Collection
}

var (
	UserAlreadyExists      = errors.New("user already exists")
	InvalidCredentials     = errors.New("invalid credentials")
	InvalidPasswordRequest = errors.New("password does not meet its requirements")
)

type User struct {
	UID      string    `json:"uid"`
	Username string    `json:"username"`
	Password string    `json:"Password"`
	Token    []*string `json:"Token"`
}

type Registration struct {
	GQLRegistration model.Registration `json:"gql_registration"`
}
