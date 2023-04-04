// Package db provides a database connection pool.
package db

import (
	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	sqlc "smart-locker/backend/db/sqlc"
)

type (
	// DB is the interface for the database querier. (Named DBTX in the official docs).
	DB interface {
		sqlc.Querier
		ExecRegisterTx(context.Context, RegisterParams) (RegisterResult, error)
		ExecLoginTx(context.Context, LoginParams) (LoginResult, error)
		ExecGetAllUserFeedsTx(context.Context, GetAllUserFeedsParams) (GetAllUserFeedsResult, error)
		ExecGetFeedByNFCSigTx(context.Context, GetFeedByNFCSigParams) (GetFeedByNFCSigResult, error)
		ExecGetAllSensorDataTx(context.Context, GetAllSensorDataParams) (GetAllSensorDataResult, error)
	}

	// Tx is the database transaction. It implements the DB interface.
	Tx struct {
		*sqlc.Queries
		db *sql.DB
	}
)

func NewTx(db *sql.DB) DB {
	return &Tx{Queries: sqlc.New(db), db: db}
}

// execTx executes a database transaction. `fn` is the bussiness logic to execute.
func (t *Tx) executeTx(ctx context.Context, fn func(*sqlc.Queries) error) error {
	tx, err := t.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := sqlc.New(tx)
	if err = fn(q); err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return rollbackErr
		}
		return err
	}
	return tx.Commit()
}
