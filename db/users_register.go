package db

import "context"

type (
	// RegisterParams is the parameters for the register query.
	RegisterParams struct {
		Email    string
		Password string
	}

	RegisterResult struct {
		Email    string
		Password string
	}
)

// ExecRegisterTx executes a transaction for the register user query.
func (*Tx) ExecRegisterTx(ctx context.Context, params RegisterParams) (RegisterResult, error) {
	panic("not implemented")
}
