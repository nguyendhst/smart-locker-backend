package api

import (
	"net/http"
	"smart-locker/backend/db"

	"github.com/labstack/echo/v4"
)

type (
	// RegisterRequest is the request body for the register user query.
	RegisterRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// RegisterResponse is the response body for the register user query.
	RegisterResponse struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
)

func (s *Server) registerUser(c echo.Context) error {
	// Get the user from the request.
	var req RegisterRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// Execute the register user query.
	params := db.RegisterParams{
		Email:    req.Email,
		Password: req.Password,
	}
	result, err := s.Store.ExecRegisterTx(c.Request().Context(), params)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	// Return the response.
	res := RegisterResponse{
		Email:    result.Email,
		Password: result.Password,
	}

	return c.JSON(http.StatusOK, res)

}
