package server

import "net/http"

type Server struct {
	router *http.ServeMux
	server *http.Server
}

func NewServer() *Server {
	router := http.NewServeMux()
	return &Server{
		router: router,
		server: &http.Server{
			Addr:    ":8080",
			Handler: router,
		},
	}
}

func (s *Server) Start() error {
	s.router.HandleFunc("/token", GetToken)
	s.router.HandleFunc("/refresh", RefreshToken)

	return s.server.ListenAndServe()
}