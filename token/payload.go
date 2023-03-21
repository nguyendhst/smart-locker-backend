package token

import (
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	uuid "github.com/google/uuid"
)

type (
	// Payload is a custom claims extending the default RegisteredClaims.
	Payload struct {
		UUID  string `json:"uuid"`
		Email string `json:"email"`
		Admin bool   `json:"admin"`
		jwt.RegisteredClaims
	}
)

func NewPayload(email string, admin bool, duration time.Duration) *Payload {
	return &Payload{
		UUID:  uuid.New().String(),
		Email: email,
		Admin: admin,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
		},
	}
}

func (p *Payload) ExpiresAt() time.Time {
	return p.RegisteredClaims.ExpiresAt.Time
}

func (p *Payload) Valid() error {
	if time.Now().After(p.ExpiresAt()) {
		return jwt.ErrInvalidKey
	}
	return nil
}
