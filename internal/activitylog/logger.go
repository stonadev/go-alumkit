package activitylog

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"
)

// Logger handles audit logging
type Logger struct {
	db *sql.DB
}

// LogEntry represents an audit log entry
type LogEntry struct {
	ID        int64
	UserID    int64
	Action    string
	Entity    string
	EntityID  int64
	Details   map[string]any
	IP        string
	UserAgent string
	CreatedAt time.Time
}

// New creates a new activity logger
func New(db *sql.DB) *Logger {
	return &Logger{db: db}
}

// Log logs an activity
func (l *Logger) Log(ctx context.Context, userID int64, action, entity string, entityID int64, details map[string]any) error {
	detailsJSON, err := json.Marshal(details)
	if err != nil {
		detailsJSON = []byte("{}")
	}

	// TODO: Insert into activity_logs table
	_ = detailsJSON
	return nil
}

// GetLogs retrieves logs for an entity
func (l *Logger) GetLogs(ctx context.Context, entity string, entityID int64) ([]LogEntry, error) {
	// TODO: Query activity_logs table
	return nil, nil
}

// GetUserLogs retrieves logs for a user
func (l *Logger) GetUserLogs(ctx context.Context, userID int64, limit int) ([]LogEntry, error) {
	// TODO: Query activity_logs table
	return nil, nil
}
