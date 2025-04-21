package main

import (
	"fmt"
	"github.com/odysseymorphey/SimpleAuth/internal/server"
	"github.com/odysseymorphey/SimpleAuth/storage/postgres"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	db, err := postgres.NewConnection(fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB")))
	if err != nil {
		logrus.Fatalf("Cannot connect to db: %v", err)
	}

	s := server.NewServer(db)

	s.Start()
}
