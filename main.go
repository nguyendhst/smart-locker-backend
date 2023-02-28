package main

import (
	"log"

	config "smart-locker/backend/config"

	echo "github.com/labstack/echo/v4"
	_ "github.com/planetscale/planetscale-go/planetscale"
)

func main() {
	// Bootstrap the application and start the server.
	e := echo.New()
	err := config.InitConfig(e)
	if err != nil {
		log.Fatal(err)
	}
	// FIXME - move this to its own file
	log.Fatal(e.Start(":" + config.PORT))
}
