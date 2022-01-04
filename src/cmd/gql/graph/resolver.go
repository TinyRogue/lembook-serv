package graph

import (
	"github.com/TinyRogue/lembook-serv/pkg/books"
	"github.com/TinyRogue/lembook-serv/pkg/user"
)

type Resolver struct {
	UserService  user.Service
	BooksService books.Service
}
