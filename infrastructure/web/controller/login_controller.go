package controller

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/xtrembaker/goflix/domain/user"
	"log"
	"net/http"
	"time"
)

const JWT_PRIVATE_KEY = "training.go"

type LoginController struct {
	UserRepository user.UserRepository
}

type LoginPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func (c LoginController) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		payload := LoginPayload{}
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			// validation error
			fmt.Fprintf(w, "Validation error")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		u, notfound := c.UserRepository.FindByUsernameAndPassword(payload.Username, payload.Password)
		if notfound != nil {
			fmt.Fprintf(w, "Invalid credentials")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"username": u.Username,
			"exp":      time.Now().Add(time.Hour * time.Duration(1)).Unix(),
			"iat":      time.Now().Unix(),
		})
		// sign token with private key
		tokenStr, err := token.SignedString([]byte(JWT_PRIVATE_KEY))
		if err != nil {
			log.Printf("Cannot generate JWT err=%v", err)
			fmt.Fprintf(w, "Error")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		JsonResponse(w, http.StatusCreated, LoginResponse{tokenStr})
	}
}
