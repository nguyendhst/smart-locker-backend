package db

import (
	"context"
	"fmt"
	"log"
	sqlc "smart-locker/backend/db/sqlc"
)

type (
	GetFeedByNFCSigParams struct {
		NFCSig string
	}

	GetFeedByNFCSigResult struct {
		Feedkey string
	}
)

func (t *Tx) ExecGetFeedByNFCSigTx(c context.Context, arg GetFeedByNFCSigParams) (GetFeedByNFCSigResult, error) {

	var res GetFeedByNFCSigResult
	var lockerId int32
	var sensorIds []int32

	err := t.executeTx(c, func(q *sqlc.Queries) error {

		// find locker by NFC signature
		locker, err := q.GetLockerByNfcSig(c, arg.NFCSig)
		if err != nil {
			log.Println("error: ", err)
			return err
		}
		lockerId = locker

		return nil
	})
	if err != nil {
		return res, err
	}

	// get sensor ids asscociated with the locker
	err = t.executeTx(c, func(q *sqlc.Queries) error {

		// find sensors for each lockers
		sensorIds, err = q.GetSensorsOfLocker(c, lockerId)
		if err != nil {
			log.Println("errorGetSensorsOfLocker: ", err)
			return err
		}
		return nil
	})
	if err != nil {
		log.Println("errorGetSensorsOfLocker: ", err)
		return res, err
	} else if len(sensorIds) == 0 {
		return res, fmt.Errorf("no sensor found for locker %v", lockerId)
	}

	// get feedkey of sensor of type lock
	log.Println("sensorIds:", sensorIds)
	for _, id := range sensorIds {
		err := t.executeTx(c, func(q *sqlc.Queries) error {

			// find sensors for each id
			log.Println("id:", id)
			sensor, err := q.GetSensorById(c, int32(id))
			if err != nil {
				log.Println(err)
			} else {

				if sensor.Kind == sqlc.SensorsKindServo {
					res.Feedkey = sensor.FeedKey
					return nil
				}
			}

			return nil
		})
		if err != nil {
			return res, err
		}
	}

	return res, nil
}
