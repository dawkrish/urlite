package main

import (
	"log"
	"net/http"
)

const (
	host = "localhost"
	port = "8080"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			page().Render(r.Context(), w)
			return
		}
		if r.Method == "POST" {
			var lightenURL string
			longURL := r.FormValue("longURL")
			if longURL == "" {
				lightenURL = "Enter a non-empty URL"
			} else {

			}
			urlResult(lightenURL).Render(r.Context(), w)
			return
		}
	})
	log.Printf("Serving at http://%v:%v\n", host, port)
	log.Fatal(http.ListenAndServe(host+":"+port, nil))
}
