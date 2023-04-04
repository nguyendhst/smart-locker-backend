package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"smart-locker/backend/api"
)

func main() {

	s, err := api.NewServer()
	if err != nil {
		log.Fatal("Server creation error: ", err)
	}
	sig := make(chan os.Signal, 1)
	// Notify the channel when an interrupt or terminate signal is received.
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	// Start the server in a goroutine so that it doesn't block.
	go func() {
		if err := api.StartServer(s); err != nil {
			log.Fatal("Server starup error: ", err)
		}
	}()

	// Wait for a signal to quit. Can be a SIGINT or SIGTERM.
	<-sig

	// Shutdown the server gracefully.
	if err := s.StopServer(); err != nil {
		log.Fatal("Server shutdown error: ", err)
	}

	log.Println("Shutting down...")

}
