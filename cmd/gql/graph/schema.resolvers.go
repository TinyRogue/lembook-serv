package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	generated2 "github.com/TinyRogue/lembook-serv/cmd/gql/graph/generated"
	model2 "github.com/TinyRogue/lembook-serv/cmd/gql/graph/generated/model"
	"github.com/TinyRogue/lembook-serv/internal/models"
	"github.com/TinyRogue/lembook-serv/pkg/middleware"
	"log"
	"net/http"
)

func (r *mutationResolver) Register(ctx context.Context, input model2.Registration) (*model2.Depiction, error) {
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

	return &model2.Depiction{
		Res: &successMsg,
	}, nil
}

func (r *mutationResolver) Login(ctx context.Context, input model2.Login) (*model2.Depiction, error) {
	w := middleware.GetResWriter(ctx)
	if w == nil {
		log.Println("Could not get writer")
		return nil, errors.New("internal server error")
	}

	log.Println("Login request from " + input.Username)
	var user models.User
	user.Username = input.Username
	user.Password = input.Password
	token, err := user.Login(ctx)
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
	return &model2.Depiction{
		Res: &msg,
	}, nil
}

func (r *queryResolver) Ping(_ context.Context) (string, error) {
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
func (r *Resolver) Mutation() generated2.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated2.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
