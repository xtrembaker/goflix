package controller

import (
	"encoding/json"
	"github.com/xtrembaker/goflix/domain/movie"
	"log"
	"net/http"
)

type MovieViewModel struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	ReleaseDate string `json:"release_date"`
	Duration    int    `json:"duration"`
	TrailerUrl  string `json:"trailer_url"`
}

func mapMovieToViewModel(m *movie.Movie) MovieViewModel {
	return MovieViewModel{
		ID:          m.ID,
		Title:       m.Title.GetValue(),
		ReleaseDate: m.ReleaseDate,
		Duration:    m.Duration,
		TrailerUrl:  m.TrailerUrl,
	}
}

func JsonResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Fatal(err)
	}
}
