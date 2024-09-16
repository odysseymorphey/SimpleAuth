package main

import (
	"github.com/odysseymorphey/SimpleAuth/internal/server"
)

func main() {
	s := server.NewServer()
	s.Start()
}