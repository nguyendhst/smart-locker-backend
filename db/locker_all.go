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

	RegisterLockerParams struct {
		UserEmail string
		NFCSig    string
	}

	RegisterLockerResult struct {
		Success bool
	}
)

func (t *Tx) ExecCreateLockerTx(c context.Context, args CreateLockerParams) (CreateLockerResult, error) {

	var res CreateLockerResult

	err := t.executeTx(c, func(q *sqlc.Queries) error {

		result, err := q.CreateLocker(c, sqlc.CreateLockerParams{
			LockerNumber: args.LockerNumber,
			Location:     args.Location,
			Status:       sqlc.LockersStatusAvailable,
			NfcSig:       args.NfcSig,
		})

		if err != nil {
			return err
		}

		//resultID, err := result.LastInsertId()
		//if err != nil {
		//	return err
		//}

		//user, err := q.GetUserByEmail(
		//	c,
		//	args.UserEmail,
		//)

		//if err != nil {
		//	return err
		//}

		//_, err = q.CreateLockerUser(c,
		//	sqlc.CreateLockerUserParams{
		//		UserID:   int32(user.ID),
		//		LockerID: int32(resultID),
		//	},
		//)
		//if err != nil {
		//	return err
		//}

		//sensorRes, err := q.CreateSensor(c,
		//	sqlc.CreateSensorParams{
		//		FeedKey: "locker" + string(rune(resultID)) + "-lock",
		//		Kind:    "servo",
		//	},
		//)

		//if err != nil {
		//	return err
		//}

		//sensorID, err := sensorRes.LastInsertId()
		//if err != nil {
		//	return err
		//}
		//_, err = q.CreateSensorLocker(
		//	c,
		//	sqlc.CreateSensorLockerParams{
		//		SensorID: int32(sensorID),
		//		LockerID: int32(resultID),
		//	},
		//)

		//if err != nil {
		//	return err
		//}

		res.Result = result
		return err
	})

	return res, err
}

func (t *Tx) ExecRegisterLockerTx(c context.Context, args RegisterLockerParams) (RegisterLockerResult, error) {
	var res RegisterLockerResult

	err := t.executeTx(c, func(q *sqlc.Queries) error {

		// find user -> get user id
		user, err := q.GetUserByEmail(
			c,
			args.UserEmail,
		)

		if err != nil {
			return err
		}

		// find locker -> get locker id
		locker, err := q.GetLockerByNfcSig(
			c,
			args.NFCSig,
		)

		if err != nil {
			return err
		}

		// create locker user
		_, err = q.CreateLockerUser(c,
			sqlc.CreateLockerUserParams{
				UserID:   int32(user.ID),
				LockerID: int32(locker.ID),
			},
		)

		if err != nil {
			return err
		}

		// update locker status
		_, err = q.UpdateLockerStatus(
			c,
			sqlc.UpdateLockerStatusParams{
				Status: sqlc.LockersStatusUsed,
				ID:     int32(locker.ID),
			},
		)

		if err != nil {
			return err
		}

		res.Success = true
		return err
	})

	return res, err
}

func (t *Tx) ExecRemoveLockerTx(c context.Context, args RemoveLockerParams) (RemoveLockerResult, error) {

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
