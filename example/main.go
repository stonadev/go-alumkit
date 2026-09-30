package main

import (
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"
	"github.com/stonadev/alumkit"
)

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func main() {
	// Load config from .env (falls back to process environment)
	cfg, err := alumkit.LoadConfig(".env")
	if err != nil {
		log.Printf("Warning: could not load .env: %v", err)
		cfg, err = alumkit.LoadConfig("")
		if err != nil {
			log.Fatalf("Failed to load config: %v", err)
		}
	}

	// Create app (connects to DB internally)
	app := alumkit.New(cfg)

	// Mount your own static files (CSS, JS, images for public pages)
	// Build with: npm run build
	app.MountStatic("/assets", os.DirFS("public"))

	// Public routes
	app.Get("/", homepageHandler(app))
	app.Get("/{slug}", pageHandler(app))

	// Add serve command and run the root command
	// (built-in: ./myapp migrate | custom: ./myapp serve)
	app.AddCommand(serveCmd(app))

	app.Execute()
}

func serveCmd(app *alumkit.App) *cobra.Command {
	var addr string
	cmd := &cobra.Command{
		Use:   "serve",
		Short: "Start the HTTP server",
		RunE: func(cmd *cobra.Command, args []string) error {
			return app.Serve(addr)
		},
	}
	cmd.Flags().StringVar(&addr, "addr", ":"+getEnv("PORT", "8080"), "Address to listen on")
	return cmd
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
