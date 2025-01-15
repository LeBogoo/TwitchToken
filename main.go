package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Options struct {
	ClientID    string `json:"clientId"`
	RedirectURI string `json:"redirectUri"`
}

func main() {
	options := Options{
		ClientID:    os.Getenv("CLIENT_ID"),
		RedirectURI: os.Getenv("REDIRECT_URI"),
	}

	http.HandleFunc("/oauth/redirect", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./public/redirect.html")
	})

	http.HandleFunc("/options", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		encoder := json.NewEncoder(w)
		encoder.Encode(options)
	})

	http.Handle("/", http.FileServer(http.Dir("./public")))

	fmt.Println("Server is listening on port 8080")
	http.ListenAndServe(":8080", nil)

}
