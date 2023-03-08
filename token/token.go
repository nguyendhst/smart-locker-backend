// Package token contains the authentication functions based on tokens.
package token

import "time"

type (
	// TokenMaker is the interface for generating and validating tokens.
	TokenMaker interface {
		// GenerateToken generates a new token for the given user.
		GenerateToken(email string, admin bool, duration time.Duration) (string, error)
		// ValidateToken validates the given token and returns the user associated with it.
		ValidateToken(token string) (*Payload, error)
	}
)
