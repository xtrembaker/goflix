package sqlite

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/xtrembaker/goflix/domain"
	"github.com/xtrembaker/goflix/domain/movie"
	"log"
	"os"
	"testing"
)

var testingDbConnection *sqlx.DB

func TestMain(m *testing.M) {
	connectTestingDb()
	ExecMigration(testingDbConnection)
	code := m.Run()
	cleanTestingDB()
	os.Exit(code)
}

func TestGetMethodReturnsEntityNotFoundWhenMovieDoesNotExists(t *testing.T) {
	movieRepository := &MovieRepository{testingDbConnection}
	m, err := movieRepository.Get(1)
	assert.Equal(t, err, errors.New(domain.EntityNotFound))
	assert.Nil(t, m, "Movie should not have been found")
}

func TestGetMethodReturnMovieWhenItExists(t *testing.T) {
	insertRecord("titre", "2018-10-18", 240, "https://www.youtube.com")
	movieRepository := &MovieRepository{testingDbConnection}
	m, err := movieRepository.Get(1)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, m, &movie.Movie{1, movie.NewTitle("titre"), "2018-10-18", 240, "https://www.youtube.com"})
}

func TestSaveMethodRecordMovieInDatabaseAndFillId(t *testing.T) {
	movieRepository := &MovieRepository{testingDbConnection}
	m := &movie.Movie{
		ID:          0,
		Title:       movie.NewTitle("Avatar: The way of Water"),
		ReleaseDate: "2022-12-25",
		Duration:    192,
		TrailerUrl:  "https://www.imdb.com/title/tt1630029/?ref_=nv_sr_srsg_0",
	}
	movieRepository.Save(m)
	assert.Greaterf(t, m.ID, int64(0), "ID should be greater than 0")
}

func connectTestingDb() {
	db, err := sqlx.Connect("sqlite3", "testing_goflix.db")
	if err != nil {
		log.Fatal(err)
	}
	testingDbConnection = db
}

func cleanTestingDB() {
	testingDbConnection.MustExec("DROP TABLE movie")
	testingDbConnection.Close()
}

func insertRecord(title string, releaseDate string, duration int, trailerUrl string) {
	testingDbConnection.MustExec(
		"INSERT INTO movie (title, release_date, duration, trailer_url) VALUES ($1, $2, $3, $4)",
		title,
		releaseDate,
		duration,
		trailerUrl,
	)
}
