package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/TinyRogue/lembook-serv/graph/generated"
	"github.com/TinyRogue/lembook-serv/graph/generated/model"
	"github.com/TinyRogue/lembook-serv/internal/models"
	"log"
)

func (r *mutationResolver) Register(ctx context.Context, input model.Registration) (*model.Depiction, error) {
	log.Println("Create new user request")
	req := models.Registration{GQLRegistration: input}

	if !models.IsPasswordValid(req.GQLRegistration.Password) {
		errMsg := models.InvalidPassword.Error()
		log.Println("Could not create user due to:" + errMsg)
		ans := model.Depiction{
			Res:   nil,
			Error: &errMsg,
		}
		return &ans, models.InvalidPassword
	}

	err := req.Save(ctx)
	if err != nil {
		errMsg := err.Error()
		log.Println("Could not create user due to:" + errMsg)
		ans := model.Depiction{
			Res:   nil,
			Error: &errMsg,
		}
		return &ans, err
	}

	successMsg := "user created"
	log.Println(successMsg)

	return &model.Depiction{
		Res:   &successMsg,
		Error: nil,
	}, nil
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	log.Println("Login request from " + input.Username)
	var user models.User
	user.Username = input.Username
	user.Password = input.Password
	//token, err := user.Login()
	//if err != nil {
	//	log.Println("Could not login user due to:" + err.Error())
	//	return "error", err
	//}

	log.Println("User logged in.")
	return "nil", nil
}

func (r *queryResolver) Ping(_ context.Context) (string, error) {
	return "Pong", nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
