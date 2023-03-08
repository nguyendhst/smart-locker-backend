package token

import (
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

type (
	JWTMaker struct {
		secret string
	}
)

func NewJWTMaker(secret string) *JWTMaker {
	return &JWTMaker{
		secret: secret,
	}
}

func (m *JWTMaker) GenerateToken(email string, admin bool, dur time.Duration) (string, error) {
	payload := NewPayload(email, admin, dur)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return token.SignedString([]byte(m.secret))

}

func (m *JWTMaker) ValidateToken(tokenString string) (*Payload, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(m.secret), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return &Payload{
			Email: claims["email"].(string),
			Admin: claims["admin"].(bool),
		}, nil
	}

	return nil, err
}
