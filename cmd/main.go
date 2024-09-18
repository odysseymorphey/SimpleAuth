package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/odysseymorphey/SimpleAuth/internal/server"
)

func main() {
	s := server.NewServer()
	s.Start()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
        <-sig
        s.Stop()
        os.Exit(1)
    }()
}