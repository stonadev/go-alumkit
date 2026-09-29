package alumkit

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
	"github.com/stonadev/alumkit/internal/auth"
	"github.com/stonadev/alumkit/internal/handler"
	"github.com/stonadev/alumkit/internal/rbac"
)

// Config holds all configuration for AlumKit
type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string

	SessionKey      string
	AppURL          string
	MaintenanceMode bool

	Mail     MailConfig
	Features FeatureConfig
	Seeder   SeederConfig
}

// MailConfig holds SMTP configuration
type MailConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	From     string
	FromName string
}

// FeatureConfig holds feature flags
type FeatureConfig struct {
	Posts     bool
	Committee bool
}

// SeederConfig holds seeder configuration
type SeederConfig struct {
	AdminName     string
	AdminEmail    string
	AdminPassword string
}

// LoadConfig loads configuration from .env file
func LoadConfig(path string) (*Config, error) {
	if path != "" {
		if err := godotenv.Load(path); err != nil {
			return nil, fmt.Errorf("failed to load .env: %w", err)
		}
	}

	mailPort, _ := strconv.Atoi(getEnv("MAIL_PORT", "587"))
	maintenanceMode, _ := strconv.ParseBool(getEnv("MAINTENANCE_MODE", "false"))
	featurePosts, _ := strconv.ParseBool(getEnv("FEATURE_POSTS", "true"))
	featureCommittee, _ := strconv.ParseBool(getEnv("FEATURE_COMMITTEE", "true"))

	cfg := &Config{
		DBHost:          getEnv("DB_HOST", ""),
		DBPort:          getEnv("DB_PORT", ""),
		DBUser:          getEnv("DB_USER", ""),
		DBPassword:      getEnv("DB_PASSWORD", ""),
		DBName:          getEnv("DB_NAME", ""),
		DBSSLMode:       getEnv("DB_SSLMODE", "disable"),
		SessionKey:      getEnv("SESSION_KEY", ""),
		AppURL:          getEnv("APP_URL", "http://localhost:8080"),
		MaintenanceMode: maintenanceMode,
		Mail: MailConfig{
			Host:     getEnv("MAIL_HOST", ""),
			Port:     mailPort,
			Username: getEnv("MAIL_USERNAME", ""),
			Password: getEnv("MAIL_PASSWORD", ""),
			From:     getEnv("MAIL_FROM", ""),
			FromName: getEnv("MAIL_FROM_NAME", "AlumKit"),
		},
		Features: FeatureConfig{
			Posts:     featurePosts,
			Committee: featureCommittee,
		},
		Seeder: SeederConfig{
			AdminName:     getEnv("ADMIN_NAME", ""),
			AdminEmail:    getEnv("ADMIN_EMAIL", ""),
			AdminPassword: getEnv("ADMIN_PASSWORD", ""),
		},
	}

	return cfg, nil
}

// App is the main application
type App struct {
	router   chi.Router
	db       *sql.DB
	session  *auth.SessionStore
	enforcer *rbac.Enforcer
	config   Config
}

// New creates a new AlumKit app
func New(cfg Config) *App {
	// Validate required fields
	if cfg.DBHost == "" {
		log.Fatal("DB_HOST is required")
	}
	if cfg.DBUser == "" {
		log.Fatal("DB_USER is required")
	}
	if cfg.DBPassword == "" {
		log.Fatal("DB_PASSWORD is required")
	}
	if cfg.DBName == "" {
		log.Fatal("DB_NAME is required")
	}
	if cfg.DBPort == "" {
		cfg.DBPort = "5432"
	}
	if cfg.DBSSLMode == "" {
		cfg.DBSSLMode = "disable"
	}

	// Build database URL from config
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName, cfg.DBSSLMode)

	// Connect to database
	db, err := sql.Open("pgx", dbURL)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	// Set connection pool settings
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)

	router := chi.NewRouter()

	// Initialize session store for user authentication
	// Uses the database to persist sessions and the session key for signing cookies
	session := auth.InitSession(db, cfg.SessionKey)
	enforcer := rbac.InitEnforcer(db)

	app := &App{
		router:   router,
		db:       db,
		session:  session,
		enforcer: enforcer,
		config:   cfg,
	}

	// Mount admin routes
	app.router.Mount("/dashboard", handler.RegisterAdminRoutes(session, enforcer, db, cfg.AppURL))

	// Mount static files
	app.router.Handle("/alumkit/*",
		http.StripPrefix("/alumkit/", http.FileServer(http.Dir("static"))))

	return app
}

// Get registers a GET route
func (a *App) Get(path string, h http.HandlerFunc) {
	a.router.Get(path, h)
}

// Post registers a POST route
func (a *App) Post(path string, h http.HandlerFunc) {
	a.router.Post(path, h)
}

// Put registers a PUT route
func (a *App) Put(path string, h http.HandlerFunc) {
	a.router.Put(path, h)
}

// Delete registers a DELETE route
func (a *App) Delete(path string, h http.HandlerFunc) {
	a.router.Delete(path, h)
}

// Route adds a route group
func (a *App) Route(path string, fn func(r chi.Router)) {
	a.router.Route(path, fn)
}

// Use adds middleware
func (a *App) Use(middleware ...func(http.Handler) http.Handler) {
	a.router.Use(middleware...)
}

// ServeHTTP implements http.Handler
func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.router.ServeHTTP(w, r)
}

// DB returns the database connection
func (a *App) DB() *sql.DB {
	return a.db
}

// Config returns the app configuration
func (a *App) Config() Config {
	return a.config
}

// CommitteeMember represents a committee member
type CommitteeMember struct {
	ID        int64
	Name      string
	PhotoPath *string
	SortOrder int
	Position  *Position
}

// Position represents a position
type Position struct {
	ID   int64
	Name string
}

// Post represents a blog post
type Post struct {
	ID          int64
	Title       string
	Body        string
	Slug        string
	Thumbnail   *string
	PublishedAt *time.Time
}

// Profile represents a user profile
type Profile struct {
	ID               int64
	UserID           int64
	PhotoPath        *string
	DateOfBirth      *time.Time
	Gender           *string
	BloodGroup       *string
	PresentAddress   *string
	PermanentAddress *string
}

// User represents a user
type User struct {
	ID        int64
	Name      string
	Email     string
	State     string
	CreatedAt time.Time
}

// RecentCommitteeMembers returns the most recent committee members
func (a *App) RecentCommitteeMembers(ctx context.Context) ([]CommitteeMember, error) {
	// TODO: Implement sqlc query
	return nil, nil
}

// PublishedPosts returns all published posts
func (a *App) PublishedPosts(ctx context.Context) ([]Post, error) {
	// TODO: Implement sqlc query
	return nil, nil
}

// PostBySlug returns a post by slug
func (a *App) PostBySlug(ctx context.Context, slug string) (*Post, error) {
	// TODO: Implement sqlc query
	return nil, nil
}

// GetProfileByUserID returns a profile by user ID
func (a *App) GetProfileByUserID(ctx context.Context, userID int64) (*Profile, error) {
	// TODO: Implement sqlc query
	return nil, nil
}

// GetUserByID returns a user by ID
func (a *App) GetUserByID(ctx context.Context, userID int64) (*User, error) {
	// TODO: Implement sqlc query
	return nil, nil
}

// RequireAuth middleware
func (a *App) RequireAuth(next http.Handler) http.Handler {
	return a.session.RequireAuth(next)
}

// RequireVerified middleware
func (a *App) RequireVerified(next http.Handler) http.Handler {
	return auth.RequireVerified(a.db, next)
}

// getEnv retrieves an environment variable with a fallback value
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
