// Package config contains the configuration for the application.
package config

import (
	"fmt"
	"os"
	"smart-locker/backend/utils"

	_ "github.com/go-sql-driver/mysql"
)

var (
	stdLogger *utils.Logger
)

// Config is the configuration struct for the application.
type Config struct {
	// The port to listen on.
	Port string `json:"port"`
	// The planetscale url
	DSN string `json:"dsn"`
	// The planetscale database
	PScaleDB string `json:"planetscale_db"`
	// Adafruit IO username
	AdafruitUsername string `json:"adafruit_username"`
	// Adafruit IO key
	AdafruitKey string `json:"adafruit_key"`
	// JWT secret
	JWTSecret string `json:"jwt_secret"`
}

func InitConfig() (*Config, error) {
	// Load the configuration file.
	var c *Config
	var err error

	if c, err = _loadConfigFile(); err != nil {
		return nil, err
	}
	return c, nil
}

func _loadConfigFile() (*Config, error) {

	var config Config = Config{}

	stdLogger = utils.NewLogger("Config")
	// Open the configuration file.
	//file, err := os.Open("config.json")
	//if err != nil {
	//	return nil, fmt.Errorf("error opening config file: %s", err)
	//}
	//defer file.Close()
	//// DEBUG -- print content of config file
	//data, _ := io.ReadAll(file)
	// log.Println(string(data))
	// Read the configuration file.
	//err = json.Unmarshal(data, &config)
	//if err != nil {
	//	return nil, fmt.Errorf("error decoding config file: %s", err)
	//}

	config.Port = os.Getenv("PORT")
	config.DSN = os.Getenv("PLANETSCALE_URL")
	config.PScaleDB = os.Getenv("PLANETSCALE_DB")
	config.AdafruitUsername = os.Getenv("ADAFRUIT_USERNAME")
	config.AdafruitKey = os.Getenv("ADAFRUIT_KEY")

	if config.Port == "" {
		config.Port = "8080"
		stdLogger.Info("No port specified in config file, defaulting to 8080")
	} else if config.DSN == "" {
		stdLogger.Info("No planetscale url specified in config file")
		return nil, fmt.Errorf("no planetscale url specified in config file")
	} else if config.PScaleDB == "" {
		stdLogger.Info("No planetscale database specified in config file")
		return nil, fmt.Errorf("no planetscale database specified in config file")
	}

	return &config, nil
}
