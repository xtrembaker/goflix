package controller

import (
	"encoding/json"
	"github.com/xtrembaker/goflix/domain/movie"
	"log"
	"net/http"
)

type createMoviePayload struct {
	Title       string `json:"title"`
	ReleaseDate string `json:"release_date"`
	Duration    int    `json:"duration"`
	TrailerUrl  string `json:"trailer_url"`
}

type CreateMovieController struct {
	MovieRepository movie.Repository
}

func (c CreateMovieController) CreateMovie() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		payload := createMoviePayload{}
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			log.Printf("Cannot parse movie body. Received=%v, Error=%v", r.Body, err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		// validation
		errorsMap := validatePayload(payload)
		if len(errorsMap) > 0 {
			JsonResponse(w, http.StatusBadRequest, errorsMap)
			return
		}

		c.MovieRepository.Save(&movie.Movie{
			ID:          0,
			Title:       movie.NewTitle(payload.Title),
			ReleaseDate: payload.ReleaseDate,
			Duration:    payload.Duration,
			TrailerUrl:  payload.TrailerUrl,
		})
		w.WriteHeader(http.StatusCreated)
	}
}

func validatePayload(p createMoviePayload) map[string]string {
	errorMap := make(map[string]string)
	if len(p.Title) <= 0 {
		errorMap["title"] = "must be at least one character"
	}
	if len(p.ReleaseDate) <= 0 {
		errorMap["release_date"] = "must be a date with format YYYY-mm-dd"
	}
	if p.Duration <= 0 {
		errorMap["duration"] = "should be at least 1 minute"
	}
	return errorMap
}
