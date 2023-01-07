package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v4"

	"hanhngo.me/m/config"
)

type JwtService struct{}

func NewJwtService() JwtService {
	return JwtService{}
}

func (*JwtService) Sign(id uint, username string) (string, error) {
	claims := jwt.MapClaims{
		"user_id":  id,
		"username": username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(config.Get("SECRET")))
	if err != nil {
		return "", err
	}

	return t, nil
}
