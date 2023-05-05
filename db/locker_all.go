package db

import (
	"context"
	sql "database/sql"
	sqlc "smart-locker/backend/db/sqlc"
)

type (
	LockerInfo struct {
	}
	//INSERT INTO lockers (locker_number, location, status, nfc_sig) VALUES (?, ?, ?, ?);
	CreateLockerParams struct {
		UserEmail    string
		LockerNumber int32
		Location     string
		//Status       sqlc.LockersStatus
		NfcSig string
	}

	CreateLockerResult struct {
		sql.Result
	}

	RemoveLockerParams struct {
		ID        int32
		UserEmail string
	}

	RemoveLockerResult struct {
		Success bool
	}
)

func (t *Tx) ExecCreateLockerTx(c context.Context, args CreateLockerParams) (CreateLockerResult, error) {

	var res CreateLockerResult

	err := t.executeTx(c, func(q *sqlc.Queries) error {

		result, err := q.CreateLocker(c, sqlc.CreateLockerParams{
			LockerNumber: args.LockerNumber,
			Location:     args.Location,
			Status:       sqlc.LockersStatusUsed,
			NfcSig:       args.NfcSig,
		})

		if err != nil {
			return err
		}

		resultID, err := result.LastInsertId()
		if err != nil {
			return err
		}

		user, err := q.GetUserByEmail(
			c,
			args.UserEmail,
		)

		if err != nil {
			return err
		}

		_, err = q.CreateLockerUser(c,
			sqlc.CreateLockerUserParams{
				UserID:   int32(user.ID),
				LockerID: int32(resultID),
			},
		)

		res.Result = result
		return err
	})

	return res, err
}

func (t *Tx) ExecRemoveLocker(c context.Context, args RemoveLockerParams) (RemoveLockerResult, error) {

	var res RemoveLockerResult

	err := t.executeTx(c, func(q *sqlc.Queries) error {

		var err error

		if err = q.DeleteLocker(c, args.ID); err != nil {
			res.Success = false
		}
		return err
	})

	return res, err
}
