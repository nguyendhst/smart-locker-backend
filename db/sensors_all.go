package db

import (
	"context"
	sqlc "smart-locker/backend/db/sqlc"
)

type (
	GetAllSensorDataParams struct{}

	GetAllSensorDataResult struct {
		Sensors []Sensor `json:"sensors"`
	}

	Sensor struct {
		ID      int64            `json:"id"`
		FeedKey string           `json:"feedkey"`
		Kind    sqlc.SensorsKind `json:"kind"`
	}
)

func (t *Tx) ExecGetAllSensorDataTx(c context.Context, arg GetAllSensorDataParams) (GetAllSensorDataResult, error) {

	var res GetAllSensorDataResult
	res.Sensors = []Sensor{}

	err := t.executeTx(c, func(q *sqlc.Queries) error {

		sensors, err := q.GetAllSensors(c)
		if err != nil {
			return err
		}

		for i := range sensors {
			res.Sensors = append(res.Sensors, Sensor{
				ID:      int64(sensors[i].ID),
				FeedKey: sensors[i].FeedKey,
				Kind:    sensors[i].Kind,
			})
		}

		return nil
	})
	if err != nil {
		return res, err
	}

	return res, nil
}
