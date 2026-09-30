package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

// RegisterAdminRoutes registers all admin routes
func (d *AdminDeps) RegisterAdminRoutes() http.Handler {
	r := chi.NewRouter()

	// Public admin routes
	r.Get("/login", d.LoginPage)
	r.Post("/login", d.Login)

	// Protected admin routes
	r.Group(func(r chi.Router) {
		r.Use(d.Session.RequireAuth)
		r.Use(d.Enforcer.RequirePermission)

		r.Get("/", d.Dashboard)
		r.Get("/logout", d.Logout)

		// User management
		r.Route("/users", func(r chi.Router) {
			r.Get("/", d.ListUsers)
			r.Get("/new", d.NewUser)
			r.Post("/new", d.CreateUser)
			r.Get("/{id}", d.EditUser)
			r.Put("/{id}", d.UpdateUser)
			r.Delete("/{id}", d.DeleteUser)
		})

		// Posts management
		r.Route("/posts", func(r chi.Router) {
			r.Get("/", d.ListPosts)
			r.Get("/new", d.NewPost)
			r.Post("/new", d.CreatePost)
			r.Get("/{id}", d.EditPost)
			r.Put("/{id}", d.UpdatePost)
			r.Delete("/{id}", d.DeletePost)
		})
	})

	return r
}
