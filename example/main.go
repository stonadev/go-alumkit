package main

import (
	"log"
	"net/http"
	"os"

	"github.com/stonadev/alumkit"
)

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func main() {
	// Load config from .env
	cfg, err := alumkit.LoadConfig(".env")
	if err != nil {
		log.Printf("Warning: could not load .env: %v", err)
		cfg = &alumkit.Config{}
	}

	// Create app (connects to DB internally)
	app := alumkit.New(*cfg)

	// Public routes
	app.Get("/", homepageHandler(app))
	app.Get("/{slug}", pageHandler(app))

	log.Printf("Listening on :%s", getEnv("PORT", "8080"))
	log.Fatal(http.ListenAndServe(":"+getEnv("PORT", "8080"), app))
}

func homepageHandler(app *alumkit.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		members, err := app.RecentCommitteeMembers(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		for _, m := range members {
			w.Write([]byte("<p>" + m.Name + "</p>"))
		}
		if len(members) == 0 {
			w.Write([]byte("<h1>Welcome to AlumKit</h1><p>No committee members found.</p>"))
		}
	}
}

func pageHandler(app *alumkit.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get slug from URL and render page
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte("<h1>Page not found</h1>"))
	}
}
