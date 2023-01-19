package in_memory

import (
	"errors"
	"github.com/xtrembaker/goflix/domain"
	"github.com/xtrembaker/goflix/domain/movie"
)

type MovieRepository struct {
	Movies []*movie.Movie
}

func (r MovieRepository) List() []*movie.Movie {
	return r.Movies
}

func (r MovieRepository) Get(id int64) (*movie.Movie, error) {
	for _, m := range r.Movies {
		if m.ID == id {
			return m, nil
		}
	}
	return nil, errors.New(domain.EntityNotFound)
}

func (r MovieRepository) Save(m *movie.Movie) {
	r.Movies = append(r.Movies, m)
}
