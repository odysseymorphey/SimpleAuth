package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/odysseymorphey/SimpleAuth/internal/server"
)

func main() {
	s := server.NewServer()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sig
		s.Stop()
		log.Println("Server stopped")
		os.Exit(0)
	}()

	err := s.Start()
	if err != nil {
		log.Fatal(err)
	}

}
