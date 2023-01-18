package web

import (
	"github.com/gorilla/mux"
	"github.com/xtrembaker/goflix/infrastructure/persistence/sqlite"
	"github.com/xtrembaker/goflix/infrastructure/web/controller"
	"net/http"
)

type Router struct {
}

func (r Router) ServeHttp() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc(
		"/",
		controller.IndexController{}.IndexRoute(),
	).Methods("GET")
	router.HandleFunc(
		"/api/movies",
		controller.MovieListController{
			MovieRepository: sqlite.MovieRepositoryFactory(),
		}.MovieList(),
	).Methods("GET")

	return router
}

func NewRouter() *Router {
	return &Router{}
}
