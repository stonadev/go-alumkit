package config

import (
	"fmt"
	"os"
	"strconv"
)

// Config holds application configuration
type Config struct {
	DB          DatabaseConfig
	Session     SessionConfig
	Mail        MailConfig
	App         AppConfig
	Features    FeatureConfig
	Seeder      SeederConfig
}

// DatabaseConfig holds database configuration
type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

// DSN returns the database connection string
func (d DatabaseConfig) DSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		d.User, d.Password, d.Host, d.Port, d.Name, d.SSLMode)
}

// SessionConfig holds session configuration
type SessionConfig struct {
	Key string
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

// AppConfig holds application configuration
type AppConfig struct {
	URL             string
	MaintenanceMode bool
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

// Load loads configuration from environment variables
func Load() *Config {
	mailPort, _ := strconv.Atoi(getEnv("MAIL_PORT", "587"))
	maintenanceMode, _ := strconv.ParseBool(getEnv("MAINTENANCE_MODE", "false"))
	featurePosts, _ := strconv.ParseBool(getEnv("FEATURE_POSTS", "true"))
	featureCommittee, _ := strconv.ParseBool(getEnv("FEATURE_COMMITTEE", "true"))

	return &Config{
		DB: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "alumkit"),
			Password: getEnv("DB_PASSWORD", ""),
			Name:     getEnv("DB_NAME", "alumkit"),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
		},
		Session: SessionConfig{
			Key: getEnv("SESSION_KEY", ""),
		},
		Mail: MailConfig{
			Host:     getEnv("MAIL_HOST", ""),
			Port:     mailPort,
			Username: getEnv("MAIL_USERNAME", ""),
			Password: getEnv("MAIL_PASSWORD", ""),
			From:     getEnv("MAIL_FROM", ""),
			FromName: getEnv("MAIL_FROM_NAME", "AlumKit"),
		},
		App: AppConfig{
			URL:             getEnv("APP_URL", "http://localhost:8080"),
			MaintenanceMode: maintenanceMode,
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
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
