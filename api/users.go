package api

import (
	"net/http"
	"smart-locker/backend/db"
	"smart-locker/backend/utils"

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
		Token string `json:"token"`
	}
)

func (s *Server) registerUser(c echo.Context) error {
	// Get the user from the request.
	var req RegisterRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	// Execute the register user query.
	params := db.RegisterParams{
		Email:          req.Email,
		PasswordHashed: hashedPassword,
	}
	result, err := s.Store.ExecRegisterTx(c.Request().Context(), params)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	// Return the response.
	res := RegisterResponse{
		Token: result.Token,
	}

	return c.JSON(http.StatusOK, res)

}
