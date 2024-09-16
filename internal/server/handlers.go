package server

import "net/http"

func GetToken(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("TOKEN"))
}

func RefreshToken(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("REFRESH"))
}