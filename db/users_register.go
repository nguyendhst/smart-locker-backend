package db

import (
	"context"
	"fmt"
	"time"

	sqlc "smart-locker/backend/db/sqlc"
	"smart-locker/backend/token"
)

var (
	TokenTimeout      = 4 * time.Hour
	ErrNoRowsAffected = fmt.Errorf("no rows affected")
)

type (
	// RegisterParams is the parameters for the register query.
	RegisterParams struct {
		Email          string
		PasswordHashed string
	}

	RegisterResult struct {
		Token string
	}
)

// ExecRegisterTx executes a transaction for the register user query.
func (t *Tx) ExecRegisterTx(ctx context.Context, params RegisterParams) (RegisterResult, error) {
	var result RegisterResult

	err := t.executeTx(ctx, func(q *sqlc.Queries) error {
		// Create the user.
		res, err := q.CreateUser(ctx, sqlc.CreateUserParams{
			Name:           "User",
			Email:          params.Email,
			PasswordHashed: params.PasswordHashed,
		})

		if err != nil {
			return err
		} else if r, _ := res.RowsAffected(); r != 1 {
			return ErrNoRowsAffected
		}

		// Create the token.
		maker := token.NewJWTMaker("secret")
		token, err := maker.GenerateToken(params.Email, false, TokenTimeout)
		if err != nil {
			return err
		}

		result = RegisterResult{
			Token: token,
		}

		return nil
	})

	return result, err
}
