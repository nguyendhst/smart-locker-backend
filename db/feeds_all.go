package db

import (
	"context"
	sqlc "smart-locker/backend/db/sqlc"
)

type (
	Locker struct {
		ID    int64  `json:"id"`
		Feeds []Feed `json:"feeds"`
	}
	Feed struct {
		ID        int64  `json:"id"`
		Feed      string `json:"feed"`
		FeedType  string `json:"feed_type"`
		FeedValue string `json:"feed_value"`
		FeedTime  string `json:"feed_time"`
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

	err := t.executeTx(c, func(q *sqlc.Queries) error {

		// find user by email
		usr, err := q.GetUserByEmail(c, arg.Email)
		if err != nil {
			return err
		}

		// find user's lockers
		lockers, err := q.GetLockersOfUser(c, usr.ID)
		if err != nil {
			return err
		}

		// find sensors for each lockers
		for i := range lockers {

			sensorIds, err := q.GetSensorsOfLocker(c, lockers[i])
			if err != nil {
				return err
			}
			res.Lockers[i] = Locker{
				ID:    int64(lockers[i]),
				Feeds: []Feed{},
			}
			for j := range sensorIds {
				feed, err := q.GetSensorById(c, sensorIds[j])
				if err != nil {
					return err
				}

				res.Lockers[i].Feeds = append(res.Lockers[i].Feeds, Feed{
					Feed:     string(feed.FeedKey),
					FeedType: string(feed.Type),
				})
			}

		}

		return err
	})

	return res, err

}
