package books

import (
	"context"
	"github.com/TinyRogue/lembook-serv/cmd/gql/graph/generated/model"
)

type Service interface {
	FindBooks(ctx context.Context, b *model.WhatBook) (model.Books, error)
}
