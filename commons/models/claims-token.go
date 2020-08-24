package models

import "github.com/dgrijalva/jwt-go"

type ClaimsToken struct {
    Username string   `json:"username"`
    Roles    []string `json:"cognito:groups"`
    jwt.StandardClaims
}
