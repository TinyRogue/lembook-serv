package books

import (
	"context"
	"github.com/TinyRogue/lembook-serv/cmd/gql/graph/generated/model"
)

func (s *Service) GetGenres(ctx context.Context, userID *string) (*model.Genres, error) {
	genres, err := getAllGenres(ctx, s)
	if err != nil {
		return nil, err
	}
	userGenres, err := getUserGenres(ctx, s, userID)
	if err != nil {
		return nil, err
	}

	for _, g := range genres.Genres {
		for _, ug := range *userGenres {
			if *ug == g.Name {
				g.Liked = true
			}
		}
	}

	return genres, nil
}
