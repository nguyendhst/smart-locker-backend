package api

import (
	"context"
	"fmt"
	"net/http"
	"smart-locker/backend/db"
	"smart-locker/backend/token"
	"sync"
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
		StartTime:  optional.NewTime(time.Now().Add(-duration).UTC()),
		EndTime:    optional.NewTime(time.Now().UTC()),
		Resolution: optional.NewInt32(int32(resolution)),
	}

	// fetch from adafruit
	wg := sync.WaitGroup{}
	for _, locker := range result.Lockers {

		fmt.Printf("Fetching %d feeds for locker %v) \n", len(locker.Feeds), locker.ID)
		for _, feed := range locker.Feeds {
			fmt.Printf("Fetching feed %v) \n", feed.Feed)
			wg.Add(1)
			ctx := context.Background()
			ctx, cancel := context.WithTimeout(ctx, time.Second*15)
			defer cancel()

			go func(f *db.Feed) {
				resp, code, err := s.AdafruitClient.DataApi.ChartData(
					ctx,
					s.Config.AdafruitUsername,
					f.Feed,
					&dataOpts,
				)
				if err != nil {
					fmt.Print("Failed fetch")
					wg.Done()
					return
				}
				fmt.Println("Fetch done")
				fmt.Println("Status code: ", code.StatusCode)
				fmt.Println("Data: ", resp.Data)
				if code.StatusCode == http.StatusOK {
					if f.FeedData == nil {
						f.FeedData = make(map[string]time.Time)
					}
					for _, values := range resp.Data {
						// convert string to time.Time
						fmt.Println("Values: ", values)
						t, err := time.Parse("2006-01-02T15:04:05Z", values[1])
						if err != nil {
							fmt.Println(err)
							continue
						}
						f.FeedData[values[0]] = t
					}
				}
				wg.Done()
			}(&feed)
		}
	}
	wg.Wait()
	// Return the response.
	res := AllFeedResponse{
		result.Lockers,
	}

	return c.JSON(http.StatusOK, res)
}
