package controller

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/xtrembaker/goflix/infrastructure/persistence/in_memory"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestItReturnsBadRequestWhenPayloadIsEmpty(t *testing.T) {
	b := bytes.Buffer{}
	json.NewEncoder(&b).Encode(&struct{}{})
	r := httptest.NewRequest("POST", "/api/movies", &b)
	w := httptest.NewRecorder()

	movieRepository := in_memory.MovieRepository{}
	CreateMovieController{MovieRepository: movieRepository}.CreateMovie()(w, r)

	var response interface{}
	assert.Equal(t, w.Result().StatusCode, http.StatusBadRequest)
	json.NewDecoder(w.Result().Body).Decode(&response)
	expectedResponse := map[string]interface{}{
		"title":        "must be at least one character",
		"release_date": "must be a date with format YYYY-mm-dd",
		"duration":     "should be at least 1 minute",
	}
	assert.Equal(t, expectedResponse, response)
}

func TestItReturnsCreatedWhenRequestHasBeenProcessed(t *testing.T) {
	b := bytes.Buffer{}
	json.NewEncoder(&b).Encode(createMoviePayload{
		Title:       "Avatar",
		ReleaseDate: "2022-12-25",
		Duration:    192,
		TrailerUrl:  "",
	})
	r := httptest.NewRequest("POST", "/api/movies", &b)
	w := httptest.NewRecorder()

	movieRepository := in_memory.MovieRepository{}
	CreateMovieController{MovieRepository: movieRepository}.CreateMovie()(w, r)

	assert.Equal(t, w.Result().StatusCode, http.StatusCreated)
}
