package handler

import (
	"github.com/stonadev/alumkit/internal/auth"
	"github.com/stonadev/alumkit/internal/rbac"
	"github.com/stonadev/alumkit/internal/service"
)

// AdminDeps holds all dependencies for admin handlers
type AdminDeps struct {
	Session  *auth.SessionStore
	Enforcer *rbac.Enforcer
	Auth     service.AuthService
	Users    service.UserService
	Posts    service.PostService
	AppURL   string
}

// New creates a new AdminDeps
func New(
	session *auth.SessionStore,
	enforcer *rbac.Enforcer,
	auth service.AuthService,
	users service.UserService,
	posts service.PostService,
	appURL string,
) *AdminDeps {
	return &AdminDeps{
		Session:  session,
		Enforcer: enforcer,
		Auth:     auth,
		Users:    users,
		Posts:    posts,
		AppURL:   appURL,
	}
}
