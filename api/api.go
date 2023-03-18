// Package api provides the API endpoints for the application.
package api

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"smart-locker/backend/config"
	"smart-locker/backend/db"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"

	swagger "github.com/nguyendhst/adafruit-go-client-v2"
)

var (
	apiEndpoints = []string{
		"/api/hello",
		"/api/users/register",
		"/api/users/login",
		"/api/feeds/all",
	}
)

type Server struct {
	Router         *echo.Echo
	Store          db.DB
	Config         *config.Config
	AdafruitClient *swagger.APIClient
}

// NewServer loads the configuration and initializes the database connection.
// It returns a new server instance.
func NewServer() (*Server, error) {

	var config *config.Config
	var db db.DB
	var err error

	if config, err = _initConfig(); err != nil {
		return nil, err
	} else if db, err = _initDB(config); err != nil {
		return nil, err
	} else if _, err = _initAdafruit(config); err != nil {
		return nil, err
	}

	e := echo.New()

	// instantiate loggin middleware
	// ripped straight from the docs
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus: true,
		LogURI:    true,
		LogMethod: true,
		BeforeNextFunc: func(c echo.Context) {
			c.Set("customValueFromContext", 42)
		},
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			//value, _ := c.Get("customValueFromContext").(int)
			fmt.Printf("REQUEST: %v, uri: %v, status: %v\n", v.Method, v.URI, v.Status)
			return nil
		},
	}))

	//e.Use(echojwt.WithConfig(echojwt.Config{
	//	SigningKey: []byte("secret"),
	//}))

	e.Logger.SetLevel(log.DEBUG)

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

	// non-restricted endpoints
	for _, endpoint := range apiEndpoints {
		switch endpoint {
		case "/api/hello":
			e.GET(endpoint, _helloWorld)
		case "/api/users/register":
			e.POST(endpoint, s.registerUser)
		case "/api/users/login":
			e.POST(endpoint, s.loginUser)
		}

	}

	// restricted endpoints
	restricted := s.Router.Group("/api/feeds")
	{
		restricted.Use(echojwt.WithConfig(echojwt.Config{
			SigningKey: []byte("secret"),
		}))

		restricted.POST("all", s.getAllFeed)

	}
	return nil
}

func _initAdafruit(config *config.Config) (*swagger.APIClient, error) {
	cfg := swagger.NewConfiguration()
	// add X-AIO-Key to header
	cfg.AddDefaultHeader("X-AIO-Key", config.AdafruitKey)
	c := swagger.NewAPIClient(cfg)
	// test ping, get all feeds
	_, _, err := c.FeedsApi.AllFeeds(context.Background(), config.AdafruitUsername)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func _helloWorld(c echo.Context) error {
	// json "message":"Hello World!"
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Hello World!",
	})
}
