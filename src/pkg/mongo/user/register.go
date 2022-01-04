package user

import (
	"context"
	"github.com/TinyRogue/lembook-serv/pkg/hash"
	nano "github.com/matoous/go-nanoid"
)

func (s *Service) Register(ctx context.Context, req *Registration) error {
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

	_, err = s.UsersCollection.InsertOne(ctx, newUser)
	return err
}
