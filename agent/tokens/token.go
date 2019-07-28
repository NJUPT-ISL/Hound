package tokens

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

// token struct
type Token struct {
	token string
}

// Generate token
func (T *Token) GenerateToken() {
	token, err := jwt.NewWithClaims(
		jwt.SigningMethodES256,
		jwt.MapClaims{"Time": time.Now()}).SigningString()
	if err != nil {
		panic(err)
	}
	T.token = token
}

// Refresh token
func (T *Token) RefreshToken() {
	T.GenerateToken()
}

// check the token
func (T *Token) VerifyToken(token string) bool {
	if token == T.token {
		return true
	} else {
		return false
	}
}

// get the token context
func (T *Token) GetToken() string {
	return T.token
}
