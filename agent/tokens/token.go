package tokens

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Token struct {
	token string
}


func (T *Token) GenerateToken() {
	token, err := jwt.NewWithClaims(
		jwt.SigningMethodES256,
		jwt.MapClaims{"Time":time.Now()}).SigningString()
	if err != nil{
		panic(err)
	}
	T.token = token
}

func (T *Token) RefreshToken() {
	T.GenerateToken()
}

func (T *Token) VerifyToken(token string) bool {
	if token == T.token{
		return true
	}else {
		return false
	}
}

func (T *Token) GetToken() string {
	return T.token
}