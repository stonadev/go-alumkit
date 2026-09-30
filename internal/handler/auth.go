package handler

import (
	"net/http"
)

// LoginPage handles GET /dashboard/login
func (d *AdminDeps) LoginPage(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Login page not implemented", http.StatusNotImplemented)
}

// Login handles POST /dashboard/login
func (d *AdminDeps) Login(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Login not implemented", http.StatusNotImplemented)
}

// Logout handles GET /dashboard/logout
func (d *AdminDeps) Logout(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Logout not implemented", http.StatusNotImplemented)
}
