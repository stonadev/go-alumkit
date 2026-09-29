package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config holds application configuration
type Config struct {
	DBHost          string
	DBPort          string
	DBUser          string
	DBPassword      string
	DBName          string
	DBSSLMode       string
	SessionKey      string
	AppURL          string
	MaintenanceMode bool
	Mail            MailConfig
	Features        FeatureConfig
	Seeder          SeederConfig
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

// DSN returns the database connection string
func (c *Config) DSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		c.DBUser, c.DBPassword, c.DBHost, c.DBPort, c.DBName, c.DBSSLMode)
}

// Load loads configuration from environment variables.
// If path is provided, it loads the .env file first.
func Load(path string) (*Config, error) {
	if path != "" {
		if err := godotenv.Load(path); err != nil {
			return nil, fmt.Errorf("failed to load .env: %w", err)
		}
	}

	mailPort, _ := strconv.Atoi(getEnv("MAIL_PORT", "587"))
	maintenanceMode, _ := strconv.ParseBool(getEnv("MAINTENANCE_MODE", "false"))
	featurePosts, _ := strconv.ParseBool(getEnv("FEATURE_POSTS", "true"))
	featureCommittee, _ := strconv.ParseBool(getEnv("FEATURE_COMMITTEE", "true"))

	return &Config{
		DBHost:          getEnv("DB_HOST", ""),
		DBPort:          getEnv("DB_PORT", ""),
		DBUser:          getEnv("DB_USER", ""),
		DBPassword:      getEnv("DB_PASSWORD", ""),
		DBName:          getEnv("DB_NAME", ""),
		DBSSLMode:       getEnv("DB_SSLMODE", "disable"),
		SessionKey:      getEnv("SESSION_KEY", ""),
		AppURL:          getEnv("APP_URL", ""),
		MaintenanceMode: maintenanceMode,
		Mail: MailConfig{
			Host:     getEnv("MAIL_HOST", ""),
			Port:     mailPort,
			Username: getEnv("MAIL_USERNAME", ""),
			Password: getEnv("MAIL_PASSWORD", ""),
			From:     getEnv("MAIL_FROM", ""),
			FromName: getEnv("MAIL_FROM_NAME", ""),
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
	}, nil
}

func getEnv(key, fallback string) string {
	if v, ok := os.LookupEnv(key); ok {
		return v
	}
	return fallback
}
