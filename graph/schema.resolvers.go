package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/TinyRogue/lembook-serv/cmd/gql/graph/generated"
	"github.com/TinyRogue/lembook-serv/cmd/gql/graph/generated/model"
)

func (r *mutationResolver) Register(ctx context.Context, input model.Registration) (*model.Depiction, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (*model.Depiction, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Ping(ctx context.Context) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) AuthorisedPing(ctx context.Context) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
