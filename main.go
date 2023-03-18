package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"smart-locker/backend/api"
)

func main() {

	sig := make(chan os.Signal, 1)
	// Notify the channel when an interrupt or terminate signal is received.
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	// Start the server in a goroutine so that it doesn't block.
	go func() {
		if err := api.StartServer(); err != nil {
			log.Fatal(err)
		}
	}()

	// Wait for a signal to quit. Can be a SIGINT or SIGTERM.
	<-sig

	log.Println("Shutting down...")

}
