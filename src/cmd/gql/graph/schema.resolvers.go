package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/TinyRogue/lembook-serv/cmd/gql/graph/generated"
	"github.com/TinyRogue/lembook-serv/cmd/gql/graph/generated/model"
	"github.com/TinyRogue/lembook-serv/pkg/middleware"
	"github.com/TinyRogue/lembook-serv/pkg/mongo/user"
)

func (r *mutationResolver) Register(ctx context.Context, input model.Registration) (*model.Depiction, error) {
	log.Println("Register new user request")
	req := user.Registration{GQLRegistration: input}

	if !user.IsPasswordValid(req.GQLRegistration.Password) {
		log.Println("Could not create user due to: " + user.InvalidPasswordRequest.Error())
		return nil, user.InvalidPasswordRequest
	}

	err := r.UserService.Register(ctx, &req)
	if err != nil {
		log.Println("Could not create user due to: " + err.Error())
		return nil, err
	}

	successMsg := "user created"
	log.Println(successMsg)

	return &model.Depiction{
		Res: &successMsg,
	}, nil
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (*model.Depiction, error) {
	w := middleware.GetResWriter(ctx)
	if w == nil {
		log.Println("Could not get writer")
		return nil, errors.New("internal server error")
	}

	log.Println("Login request from " + input.Username)
	var u user.User
	u.Username = input.Username
	u.Password = input.Password
	token, err := r.UserService.Login(ctx, &u)
	if err != nil {
		log.Println("Could not login user due to:" + err.Error())
		return nil, err
	}

	log.Println("User logged in. Setting up cookie.")
	http.SetCookie(*w, &http.Cookie{
		Name:     "auth",
		Value:    *token,
		HttpOnly: true,
		Path:     "/",
		Domain:   "localhost",
	})

	msg := "success"
	return &model.Depiction{
		Res: &msg,
	}, nil
}

func (r *queryResolver) Ping(ctx context.Context) (string, error) {
	return "Pong", nil
}

func (r *queryResolver) AuthorisedPing(ctx context.Context) (string, error) {
	u := middleware.FindUserByCtx(ctx)
	if u == nil {
		return "", fmt.Errorf("access denied")
	}
	return "Pong", nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
