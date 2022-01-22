package user

import (
	"context"
	"github.com/TinyRogue/lembook-serv/pkg/hash"
	"github.com/TinyRogue/lembook-serv/pkg/user"
	nano "github.com/matoous/go-nanoid"
)

func (s *Service) Register(ctx context.Context, req *user.Registration) error {
	if IsUsernameTaken(ctx, req.GQLRegistration.Username) {
		return user.AlreadyExists
	}

	hashedPassword, err := hash.BeautifyPassword(req.GQLRegistration.Password, nil)
	if err != nil {
		return err
	}
	UID, _ := nano.Nanoid()
	newUser := user.User{
		UID:           UID,
		Username:      req.GQLRegistration.Username,
		Password:      hashedPassword,
		Token:         []*string{},
		LikedBooks:    []*string{},
		WillingToRead: []*string{},
		DislikedBooks: []*string{},
		LikedGenres:   []*string{},
	}

	_, err = s.UsersCollection.InsertOne(ctx, newUser)
	return err
}
