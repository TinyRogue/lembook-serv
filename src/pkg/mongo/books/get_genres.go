package books

import (
	"context"
	"github.com/TinyRogue/lembook-serv/cmd/gql/graph/generated/model"
)

func (s *Service) GetGenres(ctx context.Context, userID *string) (*model.Genres, error) {
	genres, err := getAllGenres(ctx)
	if err != nil {
		return nil, err
	}
	return genres, nil
}
