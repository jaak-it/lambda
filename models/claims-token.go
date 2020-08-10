package models

import "github.com/dgrijalva/jwt-go"

type ClaimsToken struct {
	Username string `json:"cognito:username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}
