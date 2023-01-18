package controller

import (
	"fmt"
	"net/http"
)

func IndexRoute() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to Goflix !")
	}
}
