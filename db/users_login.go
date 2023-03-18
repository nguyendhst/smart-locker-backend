package db

import (
	"context"
	"fmt"

	sqlc "smart-locker/backend/db/sqlc"
	"smart-locker/backend/token"
	"smart-locker/backend/utils"
)

type (
	LoginParams struct {
		Email    string
		Password string
	}

	LoginResult struct {
		Token string
	}
)

func (t *Tx) ExecLoginTx(c context.Context, param LoginParams) (LoginResult, error) {
	var result LoginResult

	err := t.executeTx(c, func(q *sqlc.Queries) error {

		res, err := q.GetUserByEmail(c, param.Email)
		if err != nil {
			fmt.Println("query error: ", err)
			return err
		}

		if err := utils.CheckPasswordHash(param.Password, res.PasswordHashed); err != nil {
			return fmt.Errorf("passord mismatch")
		}

		// Create the token.
		maker := token.NewJWTMaker("secret")
		token, err := maker.GenerateToken(param.Email, false, TokenTimeout)
		if err != nil {
			fmt.Println("token error: ", err)
			return err
		}

		result = LoginResult{
			Token: token,
		}

		return nil
	})

	return result, err

}
