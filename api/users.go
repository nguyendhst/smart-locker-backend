package api

import (
	"fmt"
	"net/http"
	"smart-locker/backend/db"
	"smart-locker/backend/utils"

	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

var (
	// ErrUserAlreadyExists is the error returned when a user already exists.
	ErrUserAlreadyExists = echo.NewHTTPError(http.StatusBadRequest, "user already exists")
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

	// LoginRequest is the request body for the login user query.
	LoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// LoginResponse is the response body for the login user query.
	LoginResponse struct {
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
		//return c.JSON(http.StatusInternalServerError, err)
		// type assertion to check if the error is a mysql error
		mysqlErr := err.(*mysql.MySQLError)
		if err == db.ErrNoRowsAffected || mysqlErr.Number == 1062 {
			return c.JSON(http.StatusBadRequest, ErrUserAlreadyExists)
		} else {
			log.Println(err)
			return c.JSON(http.StatusInternalServerError, err)
		}
	}

	// Return the response.
	res := RegisterResponse{
		Token: result.Token,
	}
	c.Logger().Info("User registered: ", req.Email)
	return c.JSON(http.StatusCreated, res)

}

func (s *Server) loginUser(c echo.Context) error {
	// Get the user from the request.
	var req LoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// Execute the login user query.
	params := db.LoginParams{
		Email:    req.Email,
		Password: req.Password,
	}
	result, err := s.Store.ExecLoginTx(c.Request().Context(), params)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	// Return the response.
	res := LoginResponse{
		Token: result.Token,
	}
	return c.JSON(http.StatusOK, res)
}
