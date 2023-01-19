package sqlite

import (
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/xtrembaker/goflix/domain"
	"github.com/xtrembaker/goflix/domain/movie"
)

type MovieRepository struct {
	connection *sqlx.DB
}

func (r MovieRepository) List() []*movie.Movie {
	var movies []*movie.Movie
	rows, _ := r.connection.Query("SELECT * FROM movie")
	for rows.Next() {
		m := &movie.Movie{}
		var title string
		rows.Scan(&m.ID, &title, &m.ReleaseDate, &m.Duration, &m.TrailerUrl)
		m.Title = movie.NewTitle(title)
		movies = append(movies, m)
	}
	return movies
}

func (r MovieRepository) Get(id int64) (*movie.Movie, error) {
	m := &movie.Movie{}
	var title string
	row := r.connection.QueryRow("SELECT * FROM movie where id=$1", id)
	err := row.Scan(&m.ID, &title, &m.ReleaseDate, &m.Duration, &m.TrailerUrl)
	if err != nil && err == sql.ErrNoRows {
		return nil, errors.New(domain.EntityNotFound)
	}
	m.Title = movie.NewTitle(title)
	return m, nil
}

func (r MovieRepository) Save(m *movie.Movie) {
	res := r.connection.MustExec(
		"INSERT INTO movie (title, release_date, duration, trailer_url) VALUES ($1, $2, $3, $4)",
		m.Title.GetValue(),
		m.ReleaseDate,
		m.Duration,
		m.TrailerUrl,
	)
	m.ID, _ = res.LastInsertId()
}

func MovieRepositoryFactory() movie.Repository {
	var r movie.Repository = MovieRepository{
		connection: Connect(),
	}
	return r
}
