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

	uInfo := &models.UserInfo{
		GUID:      r.URL.Query().Get("GUID"),
		UserIP:    r.RemoteAddr,
		UserAgent: r.UserAgent(),
	}

	tokenPair, err := services.GeneratePair(s.db, uInfo)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

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

	uInfo := &models.UserInfo{
		GUID:      r.URL.Query().Get("GUID"),
		UserIP:    r.RemoteAddr,
		UserAgent: r.UserAgent(),
	}

	newPair, err := services.RefreshAccessToken(s.db, uInfo, tokenPair)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintln(err)))
		log.Println(err)
		return
	}

	w.Header().Set("content-type", "application/json")

	err = json.NewEncoder(w).Encode(newPair)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
}
