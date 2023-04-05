package db

import (
	"context"
	"database/sql"
	sqlc "smart-locker/backend/db/sqlc"
	"time"
)

type (
	Locker struct {
		ID           int64                  `json:"id"`
		Location     string                 `json:"location"`
		LastAccessed sql.NullTime           `json:"last_accessed,omitempty"`
		LockStatus   sqlc.LockersLockStatus `json:"lock_status"`
		Feeds        []Feed                 `json:"feeds"`
	}
	Feed struct {
		ID       int64                `json:"id"`
		Feed     string               `json:"feed"`
		FeedType string               `json:"feed_type"`
		FeedData map[time.Time]string `json:"feed_data"`
	}

	GetAllUserFeedsParams struct {
		Email string `json:"email"`
	}

	GetAllUserFeedsResult struct {
		Lockers []Locker `json:"lockers"`
	}
)

func (t *Tx) ExecGetAllUserFeedsTx(c context.Context, arg GetAllUserFeedsParams) (GetAllUserFeedsResult, error) {

	var res GetAllUserFeedsResult
	res.Lockers = []Locker{}

	err := t.executeTx(c, func(q *sqlc.Queries) error {

		// find user by email
		usr, err := q.GetUserByEmail(c, arg.Email)
		if err != nil {
			return err
		}
		// find user's lockers
		lockerIds, err := q.GetLockersOfUser(c, usr.ID)
		if err != nil {
			return err
		}

		// find sensors for each lockers
		for i := range lockerIds {

			sensorIds, err := q.GetSensorsOfLocker(c, lockerIds[i])
			//fmt.Printf("Got sensors: %v of Locker: %v\n", sensorIds, lockers[i])
			if err != nil {
				return err
			}
			locker, err := q.GetLockerByLockerNumber(c, lockerIds[i])
			if err != nil {
				return err
			}
			res.Lockers = append(res.Lockers, Locker{
				ID:           int64(lockerIds[i]),
				Location:     locker.Location,
				LastAccessed: locker.LastAccessed,
				LockStatus:   locker.LockStatus,
				Feeds:        []Feed{},
			})

			for j := range sensorIds {
				feed, err := q.GetSensorById(c, sensorIds[j])
				if err != nil {
					return err
				}

				res.Lockers[i].Feeds = append(res.Lockers[i].Feeds, Feed{
					Feed:     string(feed.FeedKey),
					FeedType: string(feed.Kind),
					FeedData: map[time.Time]string{},
				})
			}
		}

		return err
	})

	return res, err

}
