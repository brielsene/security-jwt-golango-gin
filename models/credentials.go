package models

// Create a struct to read the username and password from the request body
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Create a map to store user credentials
var Users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}
