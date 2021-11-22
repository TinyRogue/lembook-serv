package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"

	"github.com/TinyRogue/lembook-serv/graph/generated"
	"github.com/TinyRogue/lembook-serv/graph/generated/model"
	"github.com/TinyRogue/lembook-serv/internal/models"
	"github.com/TinyRogue/lembook-serv/pkg/middleware"
)

func (r *mutationResolver) Register(ctx context.Context, input model.Registration) (*model.Depiction, error) {
	log.Println("Create new user request")
	req := models.Registration{GQLRegistration: input}

	if !models.IsPasswordValid(req.GQLRegistration.Password) {
		log.Println("Could not create user due to: " + models.InvalidPasswordRequest.Error())
		return nil, models.InvalidPasswordRequest
	}

	err := req.Save(ctx)
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
	log.Println("Login request from " + input.Username)
	var user models.User
	user.Username = input.Username
	user.Password = input.Password
	token, err := user.Login(ctx)
	if err != nil {
		log.Println("Could not login user due to:" + err.Error())
		return nil, err
	}

	log.Println("User logged in.")
	return &model.Depiction{
		Res: token,
	}, nil
}

func (r *queryResolver) Ping(ctx context.Context) (string, error) {
	return "Pong", nil
}

func (r *queryResolver) AuthorisedPing(ctx context.Context) (string, error) {
	user := middleware.FindUserByCtx(ctx)
	if user == nil {
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
