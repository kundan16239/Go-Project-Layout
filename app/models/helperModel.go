package models

import "github.com/golang-jwt/jwt"

type CustomClaims struct {
	Identifier string
	jwt.StandardClaims
}
