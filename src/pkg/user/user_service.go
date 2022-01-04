package user

import (
	"context"
	"github.com/TinyRogue/lembook-serv/pkg/mongo/user"
)

type Service interface {
	Login(ctx context.Context, u *user.User) (*string, error)
	AssignNewToken(ctx context.Context, u *user.User) (*string, error)
	Register(ctx context.Context, req *user.Registration) error
}
