package server

import (
	"encoding/json"
	"fmt"
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

	tokenPair, err := services.GeneratePair(s.db, r.URL.Query().Get("GUID"), r.RemoteAddr, r.UserAgent())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.Write([]byte(r.URL.Query().Get("GUID") + "\n"))

	w.Header().Set("content-type", "application/json")

	err = json.NewEncoder(w).Encode(tokenPair)
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

	tokenPair := &models.Pair{}

	if err := json.NewDecoder(r.Body).Decode(&tokenPair); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	services.RefreshAccessToken(s.db, r.URL.Query().Get("GUID"), tokenPair)

	w.Write([]byte(fmt.Sprint(tokenPair)))

}
