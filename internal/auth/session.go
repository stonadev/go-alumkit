package auth

import (
	"context"
	"database/sql"
	"net/http"
	"time"
)

// SessionStore manages user sessions
type SessionStore struct {
	db  *sql.DB
	key string
}

type contextKey string

const userContextKey contextKey = "user"

// User represents an authenticated user
type User struct {
	ID    int64
	Name  string
	Email string
}

// InitSession creates a new session store
func InitSession(db *sql.DB, key string) *SessionStore {
	if key == "" {
		key = "default-session-key-change-me"
	}
	return &SessionStore{
		db:  db,
		key: key,
	}
}

// Create creates a new session
func (s *SessionStore) Create(w http.ResponseWriter, user User) error {
	// TODO: Implement session creation with cookie
	return nil
}

// Destroy destroys a session
func (s *SessionStore) Destroy(r *http.Request) error {
	// TODO: Implement session destruction
	return nil
}

// GetUser retrieves the user from the current session
func (s *SessionStore) GetUser(r *http.Request) (*User, error) {
	// TODO: Implement user retrieval from session
	return nil, nil
}

// RequireAuth is middleware that requires authentication
func (s *SessionStore) RequireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, err := s.GetUser(r)
		if err != nil || user == nil {
			http.Redirect(w, r, "/dashboard/login", http.StatusSeeOther)
			return
		}

		ctx := context.WithValue(r.Context(), userContextKey, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// RequireVerified is middleware that requires verified email
func RequireVerified(db *sql.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: Implement email verification check
		next.ServeHTTP(w, r)
	})
}

// GetUserFromContext retrieves user from request context
func GetUserFromContext(ctx context.Context) *User {
	user, ok := ctx.Value(userContextKey).(*User)
	if !ok {
		return nil
	}
	return user
}

// Session duration
const SessionDuration = 7 * 24 * time.Hour
