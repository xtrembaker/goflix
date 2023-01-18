package sqlite

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/xtrembaker/goflix/domain/movie"
	"log"
)

type MovieRepository struct {
	client *Client
}

func (r MovieRepository) List() []*movie.Movie {
	var movies []*movie.Movie
	err := r.client.connection.Select(&movies, "SELECT * FROM movie")
	if err != nil {
		log.Fatal(err)
	}
	return movies
}

func MovieRepositoryFactory() movie.Repository {
	var r movie.Repository = MovieRepository{
		client: getInstance(),
	}
	return r
}
