package db

import (
	"context"
	sqlc "smart-locker/backend/db/sqlc"
)

type (
	UpdateLockStatusParams struct {
		Status sqlc.LockersLockStatus
		Id     int32
	}

	UpdateLockStatusResult struct {
	}
)

func (t *Tx) ExecUpdateLockStatusTx(c context.Context, arg UpdateLockStatusParams) (UpdateLockStatusResult, error) {

	var res UpdateLockStatusResult

	err := t.executeTx(c, func(q *sqlc.Queries) error {

		_, err := q.UpdateLockStatus(c, sqlc.UpdateLockStatusParams{
			LockStatus: arg.Status,
			ID:         arg.Id,
		})

		return err
	})
	if err != nil {
		return res, err
	}

	return res, nil
}
