package utils

import "github.com/dgrijalva/jwt-go"

var JwtKey = []byte("saopaulofutebolclube")

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
