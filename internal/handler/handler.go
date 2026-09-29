package handler

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/stonadev/alumkit/internal/auth"
	"github.com/stonadev/alumkit/internal/rbac"
)

// AdminHandlers holds handler dependencies
type AdminHandlers struct {
	session  *auth.SessionStore
	enforcer *rbac.Enforcer
	db       *sql.DB
	appURL   string
}

// RegisterAdminRoutes registers admin routes
func RegisterAdminRoutes(session *auth.SessionStore, enforcer *rbac.Enforcer, db *sql.DB, appURL string) http.Handler {
	h := &AdminHandlers{
		session:  session,
		enforcer: enforcer,
		db:       db,
		appURL:   appURL,
	}

	r := chi.NewRouter()

	// Public admin routes
	r.Get("/login", h.LoginPage)
	r.Post("/login", h.Login)

	// Protected admin routes
	r.Group(func(r chi.Router) {
		r.Use(session.RequireAuth)
		r.Use(enforcer.RequirePermission)

		r.Get("/", h.Dashboard)
		r.Get("/logout", h.Logout)

		// User management
		r.Route("/users", func(r chi.Router) {
			r.Get("/", h.ListUsers)
			r.Get("/new", h.NewUser)
			r.Post("/new", h.CreateUser)
			r.Get("/{id}", h.EditUser)
			r.Put("/{id}", h.UpdateUser)
			r.Delete("/{id}", h.DeleteUser)
		})

		// Posts management
		r.Route("/posts", func(r chi.Router) {
			r.Get("/", h.ListPosts)
			r.Get("/new", h.NewPost)
			r.Post("/new", h.CreatePost)
			r.Get("/{id}", h.EditPost)
			r.Put("/{id}", h.UpdatePost)
			r.Delete("/{id}", h.DeletePost)
		})
	})

	return r
}

// LoginPage handles GET /admin/login
func (h *AdminHandlers) LoginPage(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Login page not implemented", http.StatusNotImplemented)
}

// Login handles POST /admin/login
func (h *AdminHandlers) Login(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Login not implemented", http.StatusNotImplemented)
}

// Dashboard handles GET /admin
func (h *AdminHandlers) Dashboard(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Dashboard not implemented", http.StatusNotImplemented)
}

// Logout handles GET /admin/logout
func (h *AdminHandlers) Logout(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Logout not implemented", http.StatusNotImplemented)
}

// ListUsers handles GET /admin/users
func (h *AdminHandlers) ListUsers(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "List users not implemented", http.StatusNotImplemented)
}

// NewUser handles GET /admin/users/new
func (h *AdminHandlers) NewUser(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "New user form not implemented", http.StatusNotImplemented)
}

// CreateUser handles POST /admin/users/new
func (h *AdminHandlers) CreateUser(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Create user not implemented", http.StatusNotImplemented)
}

// EditUser handles GET /admin/users/{id}
func (h *AdminHandlers) EditUser(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Edit user not implemented", http.StatusNotImplemented)
}

// UpdateUser handles PUT /admin/users/{id}
func (h *AdminHandlers) UpdateUser(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Update user not implemented", http.StatusNotImplemented)
}

// DeleteUser handles DELETE /admin/users/{id}
func (h *AdminHandlers) DeleteUser(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Delete user not implemented", http.StatusNotImplemented)
}

// ListPosts handles GET /admin/posts
func (h *AdminHandlers) ListPosts(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "List posts not implemented", http.StatusNotImplemented)
}

// NewPost handles GET /admin/posts/new
func (h *AdminHandlers) NewPost(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "New post form not implemented", http.StatusNotImplemented)
}

// CreatePost handles POST /admin/posts/new
func (h *AdminHandlers) CreatePost(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Create post not implemented", http.StatusNotImplemented)
}

// EditPost handles GET /admin/posts/{id}
func (h *AdminHandlers) EditPost(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Edit post not implemented", http.StatusNotImplemented)
}

// UpdatePost handles PUT /admin/posts/{id}
func (h *AdminHandlers) UpdatePost(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Update post not implemented", http.StatusNotImplemented)
}

// DeletePost handles DELETE /admin/posts/{id}
func (h *AdminHandlers) DeletePost(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Delete post not implemented", http.StatusNotImplemented)
}
