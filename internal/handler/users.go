package handler

import (
	"net/http"
)

// ListUsers handles GET /dashboard/users
func (d *AdminDeps) ListUsers(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "List users not implemented", http.StatusNotImplemented)
}

// NewUser handles GET /dashboard/users/new
func (d *AdminDeps) NewUser(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "New user form not implemented", http.StatusNotImplemented)
}

// CreateUser handles POST /dashboard/users/new
func (d *AdminDeps) CreateUser(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Create user not implemented", http.StatusNotImplemented)
}

// EditUser handles GET /dashboard/users/{id}
func (d *AdminDeps) EditUser(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Edit user not implemented", http.StatusNotImplemented)
}

// UpdateUser handles PUT /dashboard/users/{id}
func (d *AdminDeps) UpdateUser(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Update user not implemented", http.StatusNotImplemented)
}

// DeleteUser handles DELETE /dashboard/users/{id}
func (d *AdminDeps) DeleteUser(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Delete user not implemented", http.StatusNotImplemented)
}
