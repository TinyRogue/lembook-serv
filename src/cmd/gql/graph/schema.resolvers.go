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
	uexec "github.com/TinyRogue/lembook-serv/pkg/mongo/user"
	us "github.com/TinyRogue/lembook-serv/pkg/user"
)

func (r *mutationResolver) LikeGenre(ctx context.Context, input string) (*model.Depiction, error) {
	log.Printf("Like genre %s request.", input)
	u := middleware.FindUserByCtx(ctx)
	if u == nil {
		log.Println("Attempt to access resource without privileges. Access Denied.")
		return nil, fmt.Errorf("access denied")
	}
	err := r.BooksService.LikeGenre(ctx, &input, &u.UID)
	if err != nil {
		log.Printf("Could not like the genre due to: %v\n", err.Error())
		return nil, err
	}

	msg := "Successfully liked the genre: " + input
	log.Println(msg)
	res := model.Depiction{
		Res: &msg,
	}
	return &res, nil
}

func (r *mutationResolver) DislikeGenre(ctx context.Context, input string) (*model.Depiction, error) {
	log.Printf("Dislike genre %s request.", input)
	u := middleware.FindUserByCtx(ctx)
	if u == nil {
		log.Println("Attempt to access resource without privileges. Access Denied.")
		return nil, fmt.Errorf("access denied")
	}
	err := r.BooksService.DislikeGenre(ctx, &input, &u.UID)
	if err != nil {
		log.Printf("Could not dislike the genre due to: %v\n", err.Error())
		return nil, err
	}

	msg := "Successfully disliked the genre: " + input
	log.Println(msg)
	res := model.Depiction{
		Res: &msg,
	}
	return &res, nil
}

func (r *mutationResolver) Register(ctx context.Context, input model.Registration) (*model.Depiction, error) {
	log.Println("Register new user request")
	req := us.Registration{GQLRegistration: input}

	if !uexec.IsPasswordValid(req.GQLRegistration.Password) {
		log.Println("Could not create user due to: " + us.InvalidPasswordRequest.Error())
		return nil, us.InvalidPasswordRequest
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
	var u us.User
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

	return &model.Depiction{
		Res: &u.UID,
	}, nil
}

func (r *mutationResolver) LoginWithJwt(ctx context.Context) (*model.UserMeta, error) {
	u := middleware.FindUserByCtx(ctx)
	if u == nil {
		log.Println("Attempt to access resource without privileges. Access Denied.")
		return nil, fmt.Errorf("access denied")
	}
	return &model.UserMeta{UID: u.UID, Username: u.Username}, nil
}

func (r *queryResolver) Ping(ctx context.Context) (string, error) {
	return "Pong", nil
}

func (r *queryResolver) AuthorisedPing(ctx context.Context) (string, error) {
	u := middleware.FindUserByCtx(ctx)
	if u == nil {
		log.Println("Attempt to access resource without privileges. Access Denied.")
		return "", fmt.Errorf("access denied")
	}
	return "Pong", nil
}

func (r *queryResolver) Books(ctx context.Context, input *model.UserID) (*model.UsersBooks, error) {
	books, err := r.BooksService.FindBooks(ctx, &input.ID)
	if err != nil {
		return nil, err
	}
	return &books, nil
}

func (r *queryResolver) Genres(ctx context.Context, input *model.UserID) (*model.Genres, error) {
	log.Println("Get genres request.")
	u := middleware.FindUserByCtx(ctx)
	if u == nil {
		log.Println("Attempt to access resource without privileges. Access Denied.")
		return nil, fmt.Errorf("access denied")
	}
	genres, err := r.BooksService.GetGenres(ctx, &input.ID)
	if err != nil {
		log.Printf("Could not fetch genres due to: %v\n", err.Error())
		return nil, err
	}
	return genres, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
