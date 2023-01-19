package controller

import (
	"github.com/xtrembaker/goflix/domain/movie"
	"net/http"
)

type MovieListController struct {
	MovieRepository movie.Repository
}

func (c MovieListController) MovieList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		movies := c.MovieRepository.List()
		var resp = make([]MovieViewModel, len(movies))
		for i, movieModel := range movies {
			resp[i] = mapMovieToViewModel(movieModel)
		}
		JsonResponse(w, http.StatusOK, resp)
	}
}
