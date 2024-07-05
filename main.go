package main

import "security-jwt/routes"

// Secret key to sign the token
var jwtKey = []byte("my_secret_key")

func main() {
	routes.HandleRequests()

}
