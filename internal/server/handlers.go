package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/odysseymorphey/SimpleAuth/internal/models"
	"github.com/odysseymorphey/SimpleAuth/internal/services"
)

func (s *Server) GenerateToken(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	token_pair, err := services.GeneratePair(r.RemoteAddr, r.UserAgent())
	if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

	s.db.SaveRefreshToken(token_pair.(models.Tokens).RefreshToken)

	w.Write([]byte(r.URL.Query().Get("GUID") + "\n"))

    w.Header().Set("content-type", "application/json")

	err = json.NewEncoder(w).Encode(token_pair)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
}

func (s *Server) RefreshToken(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var tokens models.Tokens

	if err := json.NewDecoder(r.Body).Decode(&tokens); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	
}