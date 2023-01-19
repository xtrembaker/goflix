package controller

import (
	"fmt"
	"net/http"
)

type IndexController struct {
}

func (IndexController) IndexRoute() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to Goflix !")
		w.WriteHeader(http.StatusOK)
	}
}
