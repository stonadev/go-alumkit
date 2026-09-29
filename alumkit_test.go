package alumkit

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	// Write temp .env
	envContent := `DB_HOST=testhost
DB_PORT=5433
DB_USER=testuser
DB_PASSWORD=testpass
DB_NAME=testdb
DB_SSLMODE=require
SESSION_KEY=test-session-key
APP_URL=http://test:3000
FEATURE_POSTS=false
FEATURE_COMMITTEE=false
MAIL_HOST=smtp.test.com
MAIL_PORT=465
`
	tmpFile := t.TempDir() + "/.env"
	if err := os.WriteFile(tmpFile, []byte(envContent), 0644); err != nil {
		t.Fatal(err)
	}

	cfg, err := LoadConfig(tmpFile)
	if err != nil {
		t.Fatalf("LoadConfig failed: %v", err)
	}

	if cfg.DBHost != "testhost" {
		t.Errorf("DBHost = %q, want %q", cfg.DBHost, "testhost")
	}
	if cfg.DBPort != "5433" {
		t.Errorf("DBPort = %q, want %q", cfg.DBPort, "5433")
	}
	if cfg.DBUser != "testuser" {
		t.Errorf("DBUser = %q, want %q", cfg.DBUser, "testuser")
	}
	if cfg.DBPassword != "testpass" {
		t.Errorf("DBPassword = %q, want %q", cfg.DBPassword, "testpass")
	}
	if cfg.DBName != "testdb" {
		t.Errorf("DBName = %q, want %q", cfg.DBName, "testdb")
	}
	if cfg.DBSSLMode != "require" {
		t.Errorf("DBSSLMode = %q, want %q", cfg.DBSSLMode, "require")
	}
	if cfg.SessionKey != "test-session-key" {
		t.Errorf("SessionKey = %q, want %q", cfg.SessionKey, "test-session-key")
	}
	if cfg.AppURL != "http://test:3000" {
		t.Errorf("AppURL = %q, want %q", cfg.AppURL, "http://test:3000")
	}
	if cfg.Features.Posts {
		t.Error("Features.Posts should be false")
	}
	if cfg.Features.Committee {
		t.Error("Features.Committee should be false")
	}
	if cfg.Mail.Host != "smtp.test.com" {
		t.Errorf("Mail.Host = %q, want %q", cfg.Mail.Host, "smtp.test.com")
	}
	if cfg.Mail.Port != 465 {
		t.Errorf("Mail.Port = %d, want %d", cfg.Mail.Port, 465)
	}
}

func TestLoadConfigMissingFile(t *testing.T) {
	_, err := LoadConfig("/nonexistent/.env")
	if err == nil {
		t.Error("expected error for missing file")
	}
}

func TestLoadConfigEmpty(t *testing.T) {
	// Clear any existing env vars
	keys := []string{"DB_HOST", "DB_USER", "DB_PASSWORD", "DB_NAME"}
	for _, k := range keys {
		os.Unsetenv(k)
	}

	cfg, err := LoadConfig("")
	if err != nil {
		t.Fatalf("LoadConfig with empty path failed: %v", err)
	}

	// Should have empty strings (no defaults)
	if cfg.DBHost != "" {
		t.Errorf("DBHost = %q, want empty", cfg.DBHost)
	}
}

func TestNewValidatesRequired(t *testing.T) {
	// Reset env
	for _, k := range []string{"DB_HOST", "DB_USER", "DB_PASSWORD", "DB_NAME", "DB_PORT", "DB_SSLMODE"} {
		os.Unsetenv(k)
	}

	// This should fatal because required fields are empty
	// We can't easily test log.Fatal, so we skip this for now
	// In production, New() should return (*App, error) instead of using log.Fatal
	t.Skip("skip: New() uses log.Fatal, cannot test without subprocess")
}

func TestRouteRegistration(t *testing.T) {
	// Skip DB connection test
	t.Skip("requires database connection")
}

func TestAppServeHTTP(t *testing.T) {
	// Create a minimal app without DB
	app := &App{
		router: newRouter(),
	}

	handlerCalled := false
	app.Get("/test", func(w http.ResponseWriter, r *http.Request) {
		handlerCalled = true
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	req := httptest.NewRequest("GET", "/test", nil)
	w := httptest.NewRecorder()

	app.ServeHTTP(w, req)

	if !handlerCalled {
		t.Error("handler was not called")
	}
	if w.Code != http.StatusOK {
		t.Errorf("status = %d, want %d", w.Code, http.StatusOK)
	}
	if w.Body.String() != "ok" {
		t.Errorf("body = %q, want %q", w.Body.String(), "ok")
	}
}

func TestAppServeHTTPMultipleMethods(t *testing.T) {
	app := &App{
		router: newRouter(),
	}

	methods := []string{"GET", "POST", "PUT", "DELETE"}
	for _, method := range methods {
		m := method // capture for closure
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(m))
		})
		app.router.Method(m, "/"+m, handler)
	}

	for _, method := range methods {
		req := httptest.NewRequest(method, "/"+method, nil)
		w := httptest.NewRecorder()
		app.ServeHTTP(w, req)

		if w.Body.String() != method {
			t.Errorf("method %s: body = %q, want %q", method, w.Body.String(), method)
		}
	}
}
