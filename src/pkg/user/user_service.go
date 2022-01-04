package user

import (
	"context"
	"errors"
	"github.com/TinyRogue/lembook-serv/cmd/gql/graph/generated/model"
)

type Service interface {
	Login(ctx context.Context, u *User) (*string, error)
	AssignNewToken(ctx context.Context, u *User) (*string, error)
	Register(ctx context.Context, req *Registration) error
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

var (
	AlreadyExists          = errors.New("user already exists")
	InvalidCredentials     = errors.New("invalid credentials")
	InvalidPasswordRequest = errors.New("password does not meet its requirements")
)
