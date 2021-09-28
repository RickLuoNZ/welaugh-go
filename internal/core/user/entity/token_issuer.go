package entity

import (
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type TokenIssuer struct{}

type Token struct {
	Value           string
	ExpiresInSec    uint64
}

const TOKEN_EXPIRES_IN_DAY string = "3"
const TOKEN_JWT_SECRECT string = "123456asdfgh"

func (i *TokenIssuer) NewToken(id uint64) Token {
	expConfig := os.Getenv("TOKEN_EXPIRES_IN_DAY")
	if expConfig == "" {
		expConfig = TOKEN_EXPIRES_IN_DAY
	}

	expInDays, _ := strconv.ParseUint(expConfig, 10, 64)

	expInSeconds := 24 * 60 * 60 * expInDays

	secretConfig := os.Getenv("TOKEN_JWT_SECRECT")
	if secretConfig == "" {
		secretConfig = TOKEN_JWT_SECRECT
	}

	token := generateJWT(id, expInSeconds, secretConfig)

	return Token{
		Value:           token,
		ExpiresInSec:    expInSeconds,
	}
}

// create a JWT token based on id, expire date in second, and JWT secret
func generateJWT(id uint64, exp uint64, secret string) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Second * time.Duration(exp)).Unix()
	t, _ := token.SignedString([]byte(secret))
	return t
}
