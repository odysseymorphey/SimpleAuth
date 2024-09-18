package server

import (
	"log"
	"net/http"

	"github.com/odysseymorphey/SimpleAuth/internal/postgres"
)

type Server struct {
	router *http.ServeMux
	server *http.Server
	db *postgres.DB
}

func NewServer() *Server {
	router := http.NewServeMux()
	pg, err := postgres.NewConnection()
	if err != nil {
		log.Fatal(err)
	}

	return &Server{
		router: router,
		server: &http.Server{
			Addr:    ":8080",
			Handler: router,
		},
		db: pg,
	}
}

func (s *Server) Start() error {
	s.router.HandleFunc("/token", GenerateToken)
	s.router.HandleFunc("/refresh", RefreshToken)

	return s.server.ListenAndServe()
}

func (s *Server) Stop() {
	s.db.Close()
}