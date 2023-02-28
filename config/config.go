// Package config contains the configuration for the application.
package config

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"smart-locker/backend/db"
	"smart-locker/backend/utils"
	"time"

	_ "github.com/go-sql-driver/mysql"

	api "smart-locker/backend/api"

	"github.com/labstack/echo/v4"
)

var (
	PORT      string
	stdLogger *utils.Logger
	config    Config = Config{}
)

// Config is the configuration struct for the application.
type Config struct {
	// The port to listen on.
	Port string `json:"port"`
	// The planetscale url
	DSN string `json:"dsn"`
	// The planetscale database
	PScaleDB string `json:"planetscale_db"`
}

func InitConfig(e *echo.Echo) error {
	// Load the configuration file.
	var err error

	if err = _loadConfigFile(); err != nil {
		return err
	}

	// Init the DB
	if err = _initDatabase(config.DSN); err != nil {
		return err
	}

	// Init the API
	if err = _initApi(e); err != nil {
		return err
	}

	return nil
}

func _loadConfigFile() error {

	stdLogger = utils.NewLogger("Config")
	// Open the configuration file.
	file, err := os.Open("config.json")
	if err != nil {
		return fmt.Errorf("error opening config file: %s", err)
	}
	defer file.Close()
	// DEBUG -- print content of config file
	data, _ := io.ReadAll(file)
	// log.Println(string(data))
	// Read the configuration file.
	err = json.Unmarshal(data, &config)
	if err != nil {
		return fmt.Errorf("error decoding config file: %s", err)
	}

	if config.Port == "" {
		config.Port = "8080"
		stdLogger.Info("No port specified in config file, defaulting to 8080")
	} else if config.DSN == "" {
		stdLogger.Info("No planetscale url specified in config file")
		return fmt.Errorf("no planetscale url specified in config file")
	} else if config.PScaleDB == "" {
		stdLogger.Info("No planetscale database specified in config file")
		return fmt.Errorf("no planetscale database specified in config file")
	}

	PORT = config.Port

	return nil
}

func _initDatabase(token string) error {

	var err error

	// connect to the database
	if db.Conn, err = sql.Open("mysql", token); err != nil {
		return err
	}
	db.Conn.SetConnMaxLifetime(time.Minute * 1)
	db.Conn.SetMaxOpenConns(10)
	db.Conn.SetMaxIdleConns(10)

	// ping the database
	if err = db.Conn.Ping(); err != nil {
		return err
	} else {
		stdLogger.Info("Connected to database")
	}

	return nil
}

func _initApi(e *echo.Echo) error {
	// Init the API
	// Note: this is a temporary solution. API endpoints should be defined in config
	return api.InitApiHandlers(e)
}
