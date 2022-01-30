package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/TinyRogue/lembook-serv/cmd/gql/graph/generated"
	"github.com/TinyRogue/lembook-serv/cmd/gql/graph/generated/model"
	"github.com/TinyRogue/lembook-serv/pkg/middleware"
	uexec "github.com/TinyRogue/lembook-serv/pkg/mongo/user"
	us "github.com/TinyRogue/lembook-serv/pkg/user"
)

func (r *mutationResolver) LoveBook(ctx context.Context, input string) (*model.Depiction, error) {
	log.Printf("Love the book: %s request.", input)
	u := middleware.FindUserByCtx(ctx)
	if u == nil {
		log.Println("Attempt to access resource without privileges. Access Denied.")
		return nil, fmt.Errorf("access denied")
	}

	err := r.BooksService.LoveTheBook(ctx, &input, &u.UID)
	if err != nil {
		log.Printf("Could not love the book due to: %v\n", err.Error())
		return nil, err
	}

	msg := "Successfully fell in love with the book: " + input
	log.Println(msg)
	res := model.Depiction{
		Res: &msg,
	}
	return &res, nil
}

func (r *mutationResolver) DislikeBook(ctx context.Context, input string) (*model.Depiction, error) {
	log.Printf("Dislike book: %s request.", input)
	u := middleware.FindUserByCtx(ctx)
	if u == nil {
		log.Println("Attempt to access resource without privileges. Access Denied.")
		return nil, fmt.Errorf("access denied")
	}

	err := r.BooksService.DislikeTheBook(ctx, &input, &u.UID)
	if err != nil {
		log.Printf("Could not dislike the book due to: %v\n", err.Error())
		return nil, err
	}

	msg := "Successfully fulfilled the user with hatred about the book: " + input
	log.Println(msg)
	res := model.Depiction{
		Res: &msg,
	}
	return &res, nil
}

func (r *mutationResolver) AddBookToWtr(ctx context.Context, input string) (*model.Depiction, error) {
	log.Printf("Add book: %s to want to read list request.", input)
	u := middleware.FindUserByCtx(ctx)
	if u == nil {
		log.Println("Attempt to access resource without privileges. Access Denied.")
		return nil, fmt.Errorf("access denied")
	}

	err := r.BooksService.AddBookToWTR(ctx, &input, &u.UID)
	if err != nil {
		log.Printf("Could not add the book to the Want-To-Read list due to: %v\n", err.Error())
		return nil, err
	}

	msg := "Successfully added the book to the Want-To-Read list: " + input
	log.Println(msg)
	res := model.Depiction{
		Res: &msg,
	}
	return &res, nil
}

func (r *mutationResolver) CancelLoveBook(ctx context.Context, input string) (*model.Depiction, error) {
	log.Printf("Cancel love the book: %s request.", input)
	u := middleware.FindUserByCtx(ctx)
	if u == nil {
		log.Println("Attempt to access resource without privileges. Access Denied.")
		return nil, fmt.Errorf("access denied")
	}

	err := r.BooksService.CancelLoveTheBook(ctx, &input, &u.UID)
	if err != nil {
		log.Printf("Could not cancel the love to the book due to: %v\n", err.Error())
		return nil, err
	}

	msg := "Successfully cancelled love to the book: " + input
	log.Println(msg)
	res := model.Depiction{
		Res: &msg,
	}
	return &res, nil
}

func (r *mutationResolver) CancelDislikeBook(ctx context.Context, input string) (*model.Depiction, error) {
	log.Printf("Cancel dislike the book: %s request.", input)
	u := middleware.FindUserByCtx(ctx)
	if u == nil {
		log.Println("Attempt to access resource without privileges. Access Denied.")
		return nil, fmt.Errorf("access denied")
	}

	err := r.BooksService.CancelDislikeTheBook(ctx, &input, &u.UID)
	if err != nil {
		log.Printf("Could not cancel dislike the book due to: %v\n", err.Error())
		return nil, err
	}

	msg := "Successfully cancelled dislike to the book: " + input
	log.Println(msg)
	res := model.Depiction{
		Res: &msg,
	}
	return &res, nil
}

func (r *mutationResolver) CancelAddBookToWtr(ctx context.Context, input string) (*model.Depiction, error) {
	log.Printf("Cancel adding the book to Want-To-Read list: %s request.", input)
	u := middleware.FindUserByCtx(ctx)
	if u == nil {
		log.Println("Attempt to access resource without privileges. Access Denied.")
		return nil, fmt.Errorf("access denied")
	}

	err := r.BooksService.CancelAddBookToWTR(ctx, &input, &u.UID)
	if err != nil {
		log.Printf("Could not cancel adding the book to WTR list due to: %v\n", err.Error())
		return nil, err
	}

	msg := "Successfully cancelled adding the book to WTR: " + input
	log.Println(msg)
	res := model.Depiction{
		Res: &msg,
	}
	return &res, nil
}

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
	input = strings.ToLower(input)
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
		Domain:   "lembook-serv-szncc.ondigitalocean.app",
		Path:     "/",
		Secure:   true,
		SameSite: http.SameSiteNoneMode,
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

func (r *queryResolver) CategorizedBooks(ctx context.Context, input *model.UserID) (*model.UsersBooks, error) {
	log.Printf("Get books by category request.")
	u := middleware.FindUserByCtx(ctx)
	if u == nil {
		log.Println("Attempt to access resource without privileges. Access Denied.")
		return nil, fmt.Errorf("access denied")
	}
	books, err := r.BooksService.GetCategorizedBooks(ctx, &input.ID)
	if err != nil {
		log.Printf("Could not retrieve books due to: %v\n", err.Error())
		return nil, err
	}
	log.Println("Get books by category request --> success")
	return books, nil
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

func (r *queryResolver) LovedBooks(ctx context.Context) (*model.UsersBooks, error) {
	log.Println("Get loved book request.")
	u := middleware.FindUserByCtx(ctx)
	if u == nil {
		log.Println("Attempt to access resource without privileges. Access Denied.")
		return nil, fmt.Errorf("access denied")
	}

	books, err := r.BooksService.GetLovedBooks(ctx, &u.UID, 0)
	if err != nil {
		log.Printf("Could not retrieve books due to: %v\n", err.Error())
		return nil, err
	}
	log.Println("Get loved list request --> success")
	return books, nil
}

func (r *queryResolver) DislikedBooks(ctx context.Context) (*model.UsersBooks, error) {
	log.Println("Get disliked books request.")
	u := middleware.FindUserByCtx(ctx)
	if u == nil {
		log.Println("Attempt to access resource without privileges. Access Denied.")
		return nil, fmt.Errorf("access denied")
	}

	books, err := r.BooksService.GetDislikedBooks(ctx, &u.UID, 0)
	if err != nil {
		log.Printf("Could not retrieve books due to: %v\n", err.Error())
		return nil, err
	}
	log.Println("Get disliked list request --> success")
	return books, nil
}

func (r *queryResolver) WtrBooks(ctx context.Context) (*model.UsersBooks, error) {
	log.Println("Get Want-To-Read list request.")
	u := middleware.FindUserByCtx(ctx)
	if u == nil {
		log.Println("Attempt to access resource without privileges. Access Denied.")
		return nil, fmt.Errorf("access denied")
	}

	books, err := r.BooksService.GetWTRBooks(ctx, &u.UID, 0)
	if err != nil {
		log.Printf("Could not retrieve books due to: %v\n", err.Error())
		return nil, err
	}
	log.Println("Get Want-To-Read list request --> success")
	return books, nil
}

func (r *queryResolver) PredictedBooks(ctx context.Context) (*model.UsersBooks, error) {
	log.Println("Get predicted list request.")
	u := middleware.FindUserByCtx(ctx)
	if u == nil {
		log.Println("Attempt to access resource without privileges. Access Denied.")
		return nil, fmt.Errorf("access denied")
	}

	books, err := r.BooksService.GetPredictedBooks(ctx, &u.UID, 28)
	if err != nil {
		log.Printf("Could not retrieve books due to: %v\n", err.Error())
		return nil, err
	}
	log.Println("Get predicted list request --> success")
	return books, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
