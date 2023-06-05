package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// Keep our string in memory
var url string = ""

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	// Redirect to the URL
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		if url == "" {
			w.Write([]byte("No URL is currently set"))
			return
		}

		http.Redirect(w, r, url, http.StatusFound)
	})

	// See the url
	r.Get("/see", func(w http.ResponseWriter, r *http.Request) {
		if url == "" {
			w.Write([]byte("No URL is currently set"))
			return
		}

		w.Write([]byte(url))
	})

	// Set the URL
	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		// Get the text from the body
		body, err := io.ReadAll(r.Body)
		if err != nil {
			w.Write([]byte("Could not set URL"))
			w.WriteHeader(400)
			return
		}

		url = string(body)
		w.Write([]byte("URL has been set to " + url))
	})

	log.Printf("Starting server on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
