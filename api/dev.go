package api

import (
	"context"
	"fmt"
	"net/http"
	"smart-locker/backend/alert"
	"smart-locker/backend/fcm"

	"github.com/labstack/echo/v4"
)

type (
	FCMPingRequest struct {
	}

	FCMPingResponse struct {
		Status string `json:"status"`
	}
)

func (s *Server) fcmPing(c echo.Context) error {
	// Get the user from the request.
	var req FCMPingRequest
	if err := c.Bind(&req); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, err)
	}

	//a, err := alert.NewAlert()
	err := alert.NewAlert()

	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	ctx := context.Background()
	client, err := alert.Alerter.FirebaseApp.Messaging(ctx)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	fcm.SendAlert(ctx, client, "alert", "test")

	return c.JSON(http.StatusOK, FCMPingResponse{Status: "OK"})
}
