package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/dharnnie/booksapi/app/types"
	"github.com/gorilla/context"
)

// CreateToken creates new token for users on login
func CreateToken(w http.ResponseWriter, r *http.Request) {
	var u types.User
	_ = json.NewDecoder(r.Body).Decode(&u)
	// never store password in token!
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"username": u.Username,
		"password": u.Password,
	})
	// secret argument should be a special and unique string
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		fmt.Println(err)
	}
	json.NewEncoder(w).Encode(types.JwtToken{Token: tokenString})
}

// NB: This wrapper function (Authenticate) was inspired by Nic Raboy's post on polyglot developer
// You can find it here: https://goo.gl/RmegnT

// Authenticate will serve as a wrapper around every http Handler for each endpoint
func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	// return a HandlerFunc that performs authourisation
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// get token as an authorisation header from request header
		authorisationHeader := r.Header.Get("authorization")
		// make sure authorisation header is not empty
		if authorisationHeader != "" {
			// We only need to check that authorisation header is not empty
			token, err := jwt.Parse(authorisationHeader, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}
				return []byte("secret"), nil
			})
			if err != nil {
				json.NewEncoder(w).Encode(types.Exception{Message: error.Error(err)})
				return
			}
			if token.Valid {
				context.Set(r, "decoded", token.Claims)
				next(w, r)
			} else {
				json.NewEncoder(w).Encode(types.Exception{Message: "Invalid authorization token"})
			}
		} else {
			json.NewEncoder(w).Encode(types.Exception{Message: "An authorization header is required"})
		}
	})
}
