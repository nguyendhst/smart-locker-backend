// Package api provides the API endpoints for the application.
package api

import (
	"database/sql"
	"net/http"
	"smart-locker/backend/config"
	"smart-locker/backend/db"

	"github.com/labstack/echo/v4"
)

type ()

var (
	apiEndpoints = []string{
		"api/hello",
		"api/users/register",
	}
)

type Server struct {
	Router *echo.Echo
	Store  db.DB
	Config *config.Config
}

func NewServer() (*Server, error) {

	var config *config.Config
	var db db.DB
	var err error

	if config, err = _initConfig(); err != nil {
		return nil, err
	} else if db, err = _initDB(config); err != nil {
		return nil, err
	}

	e := echo.New()

	return &Server{
		Router: e,
		Store:  db,
		Config: config,
	}, nil

}

func StartServer() error {

	s, err := NewServer()
	if err != nil {
		return err
	}

	if err = _initApi(s, s.Router); err != nil {
		return err
	}

	return s.Router.Start(":" + s.Config.Port)
}

func _initConfig() (*config.Config, error) {
	return config.InitConfig()
}

func _initDB(config *config.Config) (db.DB, error) {

	dbConn, err := sql.Open("mysql", config.DSN)
	if err != nil {
		return nil, err
	}

	tx := db.NewTx(
		dbConn,
	)

	return tx, nil
}

func _initApi(s *Server, e *echo.Echo) error {
	for _, endpoint := range apiEndpoints {
		switch endpoint {
		case "api/hello":
			e.GET(endpoint, _helloWorld)
		case "api/users/register":
			e.POST(endpoint, s.registerUser)
		}
	}
	return nil
}

func _helloWorld(c echo.Context) error {
	// json "message":"Hello World!"
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Hello World!",
	})
}
