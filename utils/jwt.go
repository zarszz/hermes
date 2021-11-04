package utils

import (
	"errors"
	"os"

	"github.com/golang-jwt/jwt"
)

type JWTService struct {
	JWTSecret []byte
}

func NewJWTService() *JWTService {
	return &JWTService{JWTSecret: []byte(os.Getenv("JWT_SECRET"))}
}

func (s *JWTService) Generate(userId string) (*string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = userId

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signed, err := token.SignedString(s.JWTSecret)
	if err != nil {
		return nil, err
	}
	return &signed, nil
}

func (s *JWTService) Validate(encodedToken string)(*jwt.Token, error)  {
	token, err := jwt.Parse(encodedToken, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token")
		}
		return s.JWTSecret, nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}
