package controller

import (
	"github.com/xtrembaker/goflix/domain/movie"
	"net/http"
)

type MovieListViewModel struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	ReleaseDate string `json:"release_date"`
	Duration    int    `json:"duration"`
	TrailerUrl  string `json:"trailer_url"`
}

type MovieListController struct {
	MovieRepository movie.Repository
}

func (c MovieListController) MovieList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		movies := c.MovieRepository.List()
		var resp = make([]MovieListViewModel, len(movies))
		for i, movieModel := range movies {
			resp[i] = mapMovieToViewModel(movieModel)
		}
		JsonResponse(w, resp)
	}
}

func mapMovieToViewModel(m *movie.Movie) MovieListViewModel {
	return MovieListViewModel{
		ID:          m.ID,
		Title:       m.Title,
		ReleaseDate: m.ReleaseDate,
		Duration:    m.Duration,
		TrailerUrl:  m.TrailerUrl,
	}
}
