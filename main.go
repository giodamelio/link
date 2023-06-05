package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/redis/go-redis/v9"
)

// Keep our string in memory
var URL_KEY = "URL"

func redisConnect() *redis.Client {
	redisURL := os.Getenv("REDIS_URL")
	if redisURL == "" {
		panic("REDIS_URL environment variable not set")
	}

	options, err := redis.ParseURL(redisURL)
	if err != nil {
		panic(err)
	}

	return redis.NewClient(options)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Connect to Redis Database
	rdb := redisConnect()

	// Setup Chi router
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Redirect to the URL
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		url, err := rdb.Get(context.Background(), URL_KEY).Result()
		if err != nil {
			w.Write([]byte("No URL is currently set"))
			return
		}

		http.Redirect(w, r, url, http.StatusFound)
	})

	// See the url
	r.Get("/see", func(w http.ResponseWriter, r *http.Request) {
		url, err := rdb.Get(context.Background(), URL_KEY).Result()
		if err != nil {
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
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Write the URL to Redis
		err = rdb.Set(context.Background(), URL_KEY, body, time.Minute*5).Err()
		if err != nil {
			w.Write([]byte("No URL is currently set"))
			return
		}

		w.Write([]byte("URL has been set to "))
		w.Write(body)
	})

	log.Printf("Starting server on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
