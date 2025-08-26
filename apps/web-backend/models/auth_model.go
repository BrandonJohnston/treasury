package models

// LoginRequest represents the structure of the login request
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// type LoginResponse struct {
// 	Token string `json:"token"`
// }

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
