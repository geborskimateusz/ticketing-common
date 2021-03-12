package common 

import 	"github.com/dgrijalva/jwt-go"

// UserPayload is metadata about current user in session
type UserPayload struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

type CustomClaims struct {
	Email string `json:"email"`
	ID    string `json:"id"`
	jwt.StandardClaims
}

