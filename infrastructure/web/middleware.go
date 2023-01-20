package web

import (
	jwtMiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/form3tech-oss/jwt-go"
	"github.com/xtrembaker/goflix/infrastructure/web/controller"
	"log"
	"net/http"
)

func logRequestMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%v], %v", r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	}
}

func loggedOnly(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		j := jwtMiddleware.New(jwtMiddleware.Options{
			ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
				return []byte(controller.JWT_PRIVATE_KEY), nil
			},
			SigningMethod: jwt.SigningMethodHS256,
		})
		j.HandlerWithNext(w, r, next)
	}
}
