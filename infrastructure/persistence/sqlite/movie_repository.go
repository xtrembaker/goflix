package sqlite

import (
	"errors"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/xtrembaker/goflix/domain"
	"github.com/xtrembaker/goflix/domain/movie"
	"log"
)

type MovieRepository struct {
	connection *sqlx.DB
}

func (r MovieRepository) List() []*movie.Movie {
	var movies []*movie.Movie
	err := r.connection.Select(&movies, "SELECT * FROM movie")
	if err != nil {
		log.Fatal(err)
	}
	return movies
}

func (r MovieRepository) Get(id int64) (*movie.Movie, error) {
	m := &movie.Movie{}
	err := r.connection.Get(m, "SELECT * FROM movie where id=$1", id)
	if err != nil {
		return nil, errors.New(domain.EntityNotFound) // consider any error as movie does not exist
	}
	return m, nil
}

func MovieRepositoryFactory() movie.Repository {
	var r movie.Repository = MovieRepository{
		connection: Connect(),
	}
	return r
}
