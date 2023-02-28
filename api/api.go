// Package api provides the API endpoints for the application.
package api

import (
	"net/http"
	"smart-locker/backend/utils"

	"github.com/labstack/echo/v4"
)

type ()

var (
	stdLogger *utils.Logger

	apiEndpoints = []string{
		"api/hello",
	}
)

func InitApiHandlers(e *echo.Echo) error {
	// Init the logger.
	stdLogger = utils.NewLogger("API")

	// Init the API endpoints.
	for _, endpoint := range apiEndpoints {
		switch endpoint {
		case "api/hello":
			e.GET(endpoint, helloWorld)
			stdLogger.Info("Registered endpoint: " + endpoint)
		}
	}
	return nil
}

func helloWorld(c echo.Context) error {
	// json "message":"Hello World!"
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Hello World!",
	})
}
