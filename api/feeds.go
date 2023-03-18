package api

import (
	"fmt"
	"net/http"
	"smart-locker/backend/db"

	"github.com/labstack/echo/v4"
)

type (
	AllFeedRequest struct {
		Token string `json:"token"`
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
	params := db.GetAllUserFeedsParams{
		Email: c.Get("email").(string),
	}

	result, err := s.Store.ExecGetAllUserFeedsTx(c.Request().Context(), params)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	// Return the response.
	res := AllFeedResponse{
		result.Lockers,
	}

	return c.JSON(http.StatusOK, res)
}
