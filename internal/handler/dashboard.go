package handler

import (
	"net/http"
)

// Dashboard handles GET /dashboard
func (d *AdminDeps) Dashboard(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Dashboard not implemented", http.StatusNotImplemented)
}
