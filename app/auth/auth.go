package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/dharnnie/booksapi/app/types"
)

// CreateToken creates new token for users on login
func CreateToken(w http.ResponseWriter, r *http.Request) {
	var u types.User
	_ = json.NewDecoder(r.Body).Decode(&u)
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
