package sqlite

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/xtrembaker/goflix/domain/movie"
)

type MovieRepository struct {
	db *Client
}

func MovieRepositoryFactory() movie.Repository {
	var r movie.Repository = MovieRepository{
		db: GetInstance(),
	}
	return r
}
