package auth

import "github.com/dgrijalva/jwt-go"

// NewUserData represents all the information needed to create a new user
type NewUserData struct {
	Mail     string `json:"email" binding:"required"`
	Username string `json:"username"`
	Password string `json:"password" binding:"required"`
}

// LoginCredentials represents the information needed to log in
type LoginCredentials struct {
	Mail     string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// UserClaims stores the standard user claims
type UserClaims struct {
	Mail string `json:"email" binding:"required"`
	jwt.StandardClaims
}
