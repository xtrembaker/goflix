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
		loggedOnly(
			controller.MovieListController{
				MovieRepository: sqlite.MovieRepositoryFactory(),
			}.MovieList(),
		),
	).Methods("GET")
	router.HandleFunc(
		"/api/movie/{id:[0-9]+}",
		loggedOnly(
			controller.ShowMovieController{
				MovieRepository: sqlite.MovieRepositoryFactory(),
			}.ShowMovie(),
		),
	).Methods("GET")
	router.HandleFunc(
		"/api/movie",
		loggedOnly(
			controller.CreateMovieController{
				MovieRepository: sqlite.MovieRepositoryFactory(),
			}.CreateMovie(),
		),
	).Methods("POST")
	router.HandleFunc(
		"/api/login",
		controller.LoginController{
			UserRepository: sqlite.UserRepositoryFactory(),
		}.Login(),
	).Methods("POST")

	// add middleware
	return logRequestMiddleware(router.ServeHTTP)
}

func NewRouter() *Router {
	return &Router{}
}
