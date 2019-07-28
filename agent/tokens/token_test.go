package tokens

import (
	"testing"
	"time"
)

func TestToken_GenerateToken(t *testing.T) {
	token := Token{}
	token.GenerateToken()
	t.Logf(token.GetToken())
}

func TestToken_RefreshToken(t *testing.T) {
	token := Token{}
	token.GenerateToken()
	t.Logf("The original token:" + token.GetToken())
	time.Sleep(100000)
	token.RefreshToken()
	t.Logf("The   New    token:" + token.GetToken())
}

func TestToken_VerifyToken(t *testing.T) {
	token1 := Token{}
	token1.GenerateToken()
	t.Logf("The original token1:" + token1.GetToken())
	time.Sleep(100000)
	token2 := Token{}
	token2.GenerateToken()
	t.Logf("The original token2:" + token2.GetToken())
	t.Logf("The token1 equal to to token2:" + func() string {
		if token1.VerifyToken(token2.GetToken()) {
			return "Ture"
		}
		return "False"
	}())
}
