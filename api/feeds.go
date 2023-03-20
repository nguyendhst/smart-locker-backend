package api

import (
	"context"
	"fmt"
	"net/http"
	"smart-locker/backend/db"
	"smart-locker/backend/token"
	"time"

	"github.com/antihax/optional"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	swagger "github.com/nguyendhst/adafruit-go-client-v2"
)

var (
	// 10 mins
	duration time.Duration = 10 * time.Minute
	// size of aggregation slice. Must be one of: 1, 5, 10, 30, 60, 120, 240, 480, or 960
	resolution int32 = 1
	// aggregate field. Must be one of: avg, sum, val, min, max, val_count
	//field string = "max"
	// ISO8601 formatted dates.
	endTime   time.Time = time.Now().UTC()
	startTime time.Time = time.Now().Add(-duration).UTC()
)

type (
	AllFeedRequest struct {
	}

	AllFeedResponse struct {
		Lockers []db.Locker `json:"lockers"`
	}
)

// getAllFeed returns all feed for the current user ID.
func (s *Server) getAllFeed(c echo.Context) error {
	var req AllFeedRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	//if err := c.Validate(req); err != nil {
	//	return c.JSON(http.StatusBadRequest, err)
	//}

	// Execute the query.
	// Get the email from the set context from jwt middleware
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*token.Payload)
	email := claims.Email

	params := db.GetAllUserFeedsParams{
		Email: email,
	}

	result, err := s.Store.ExecGetAllUserFeedsTx(c.Request().Context(), params)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	dataOpts := swagger.DataApiChartDataOpts{
		StartTime:  optional.NewTime(startTime),
		EndTime:    optional.NewTime(endTime),
		Resolution: optional.NewInt32(int32(resolution)),
	}

	// fetch from adafruit
	for _, locker := range result.Lockers {
		for _, feed := range locker.Feeds {
			ctx := context.Background()
			ctx, cancel := context.WithTimeout(ctx, time.Second*15)
			defer cancel()
			go func(f db.Feed) {
				s.AdafruitClient.DataApi.ChartData(
					ctx,
					s.Config.AdafruitUsername,
					f.Feed,
					&dataOpts,
				)

			}(feed)
		}
	}

	// Return the response.
	res := AllFeedResponse{
		result.Lockers,
	}

	return c.JSON(http.StatusOK, res)
}
