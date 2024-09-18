package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/odysseymorphey/SimpleAuth/internal/services"
)

func GenerateToken(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	token_pair, err := services.GeneratePair(r.RemoteAddr, r.UserAgent())
	if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

	w.Write([]byte(r.URL.Query().Get("GUID") + "\n"))

    w.Header().Set("content-type", "application/json")

	err = json.NewEncoder(w).Encode(token_pair)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
}

func RefreshToken(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("REFRESH"))
}