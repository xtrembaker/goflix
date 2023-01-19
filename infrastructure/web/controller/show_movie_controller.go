package controller

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/xtrembaker/goflix/domain"
	"github.com/xtrembaker/goflix/domain/movie"
	"net/http"
	"strconv"
)

type ShowMovieController struct {
	MovieRepository movie.Repository
}

func (c ShowMovieController) ShowMovie() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		m, err := c.MovieRepository.Get(id)
		if err != nil && err.Error() == domain.EntityNotFound {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, http.StatusText(http.StatusNotFound))
			return
		}
		JsonResponse(w, http.StatusOK, mapMovieToViewModel(m))
	}
}
