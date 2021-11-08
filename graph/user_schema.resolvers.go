package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"

	"github.com/TinyRogue/lembook-serv/graph/generated"
	"github.com/TinyRogue/lembook-serv/graph/generated/model"
	"github.com/TinyRogue/lembook-serv/internal/users"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (string, error) {
	log.Println("Create new user request")
	var user users.User
	user.Username = input.Username
	user.Password = input.Password
	err := user.Create()
	if err != nil {
		log.Println("Could not create user due to:" + err.Error())
		return "error", err
	}

	log.Println("User created.")
	return "success", nil
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	log.Println("Login request from " + input.Username)
	var user users.User
	user.Username = input.Username
	user.Password = input.Password
	token, err := user.Login()
	if err != nil {
		log.Println("Could not login user due to:" + err.Error())
		return "error", err
	}

	log.Println("User logged in.")
	return *token, nil
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Ping(ctx context.Context) (string, error) {
	return "Pong", nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
