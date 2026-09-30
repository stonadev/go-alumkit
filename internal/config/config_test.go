package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoad(t *testing.T) {
	tests := []struct {
		name    string
		env     map[string]string
		want    *Config
		wantErr bool
	}{
		{
			name: "all values set",
			env: map[string]string{
				"DB_HOST":     "db.example.com",
				"DB_PORT":     "5433",
				"DB_USER":     "myuser",
				"DB_PASSWORD": "secret",
				"DB_NAME":     "mydb",
				"DB_SSLMODE":  "require",
				"SESSION_KEY": "key123",
				"APP_URL":     "https://example.com",
				"MAIL_HOST":   "smtp.example.com",
				"MAIL_PORT":   "465",
			},
			want: &Config{
				DBHost:     "db.example.com",
				DBPort:     "5433",
				DBUser:     "myuser",
				DBPassword: "secret",
				DBName:     "mydb",
				DBSSLMode:  "require",
				SessionKey: "key123",
				AppURL:     "https://example.com",
				Mail: MailConfig{
					Host: "smtp.example.com",
					Port: 465,
				},
				Features: FeatureConfig{
					Posts:     true,
					Committee: true,
				},
			},
		},
		{
			name: "empty values with defaults",
			env:  map[string]string{},
			want: &Config{
				DBSSLMode: "disable",
				Mail: MailConfig{
					Port: 587,
				},
				Features: FeatureConfig{
					Posts:     true,
					Committee: true,
				},
			},
		},
		{
			name: "features disabled",
			env: map[string]string{
				"FEATURE_POSTS":     "false",
				"FEATURE_COMMITTEE": "false",
			},
			want: &Config{
				DBSSLMode: "disable",
				Mail: MailConfig{
					Port: 587,
				},
				Features: FeatureConfig{
					Posts:     false,
					Committee: false,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Clear env
			clearEnv(t)

			// Set test env
			for k, v := range tt.env {
				t.Setenv(k, v)
			}

			got, err := Load("")
			if (err != nil) != tt.wantErr {
				t.Fatalf("Load() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}

			if got.DBHost != tt.want.DBHost {
				t.Errorf("DBHost = %q, want %q", got.DBHost, tt.want.DBHost)
			}
			if got.DBPort != tt.want.DBPort {
				t.Errorf("DBPort = %q, want %q", got.DBPort, tt.want.DBPort)
			}
			if got.DBUser != tt.want.DBUser {
				t.Errorf("DBUser = %q, want %q", got.DBUser, tt.want.DBUser)
			}
			if got.DBPassword != tt.want.DBPassword {
				t.Errorf("DBPassword = %q, want %q", got.DBPassword, tt.want.DBPassword)
			}
			if got.DBName != tt.want.DBName {
				t.Errorf("DBName = %q, want %q", got.DBName, tt.want.DBName)
			}
			if got.DBSSLMode != tt.want.DBSSLMode {
				t.Errorf("DBSSLMode = %q, want %q", got.DBSSLMode, tt.want.DBSSLMode)
			}
			if got.SessionKey != tt.want.SessionKey {
				t.Errorf("SessionKey = %q, want %q", got.SessionKey, tt.want.SessionKey)
			}
			if got.AppURL != tt.want.AppURL {
				t.Errorf("AppURL = %q, want %q", got.AppURL, tt.want.AppURL)
			}
			if got.Mail.Host != tt.want.Mail.Host {
				t.Errorf("Mail.Host = %q, want %q", got.Mail.Host, tt.want.Mail.Host)
			}
			if got.Mail.Port != tt.want.Mail.Port {
				t.Errorf("Mail.Port = %d, want %d", got.Mail.Port, tt.want.Mail.Port)
			}
			if got.Features.Posts != tt.want.Features.Posts {
				t.Errorf("Features.Posts = %v, want %v", got.Features.Posts, tt.want.Features.Posts)
			}
			if got.Features.Committee != tt.want.Features.Committee {
				t.Errorf("Features.Committee = %v, want %v", got.Features.Committee, tt.want.Features.Committee)
			}
		})
	}
}

func TestLoadFromDotEnv(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	envFile := filepath.Join(dir, ".env")

	content := `DB_HOST=from-file
DB_USER=fileuser
DB_PASSWORD=filepass
DB_NAME=filedb
`
	if err := os.WriteFile(envFile, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	clearEnv(t)

	got, err := Load(envFile)
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}

	if got.DBHost != "from-file" {
		t.Errorf("DBHost = %q, want %q", got.DBHost, "from-file")
	}
	if got.DBUser != "fileuser" {
		t.Errorf("DBUser = %q, want %q", got.DBUser, "fileuser")
	}
}

func TestLoadFromDotEnvMissing(t *testing.T) {
	t.Parallel()

	_, err := Load("/nonexistent/.env")
	if err == nil {
		t.Error("expected error for missing .env file")
	}
}

func TestDSN(t *testing.T) {
	t.Parallel()

	cfg := &Config{
		DBHost:     "localhost",
		DBPort:     "5432",
		DBUser:     "user",
		DBPassword: "pass",
		DBName:     "db",
		DBSSLMode:  "disable",
	}

	want := "postgres://user:pass@localhost:5432/db?sslmode=disable"
	got := cfg.DSN()

	if got != want {
		t.Errorf("DSN() = %q, want %q", got, want)
	}
}

func clearEnv(t *testing.T) {
	t.Helper()
	keys := []string{
		"DB_HOST", "DB_PORT", "DB_USER", "DB_PASSWORD", "DB_NAME", "DB_SSLMODE",
		"SESSION_KEY", "APP_URL", "MAINTENANCE_MODE",
		"FEATURE_POSTS", "FEATURE_COMMITTEE",
		"MAIL_HOST", "MAIL_PORT", "MAIL_USERNAME", "MAIL_PASSWORD", "MAIL_FROM", "MAIL_FROM_NAME",
	}
	for _, k := range keys {
		os.Unsetenv(k)
	}
}
