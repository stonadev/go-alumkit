package handler

import (
	"net/http"
)

// ListPosts handles GET /dashboard/posts
func (d *AdminDeps) ListPosts(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "List posts not implemented", http.StatusNotImplemented)
}

// NewPost handles GET /dashboard/posts/new
func (d *AdminDeps) NewPost(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "New post form not implemented", http.StatusNotImplemented)
}

// CreatePost handles POST /dashboard/posts/new
func (d *AdminDeps) CreatePost(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Create post not implemented", http.StatusNotImplemented)
}

// EditPost handles GET /dashboard/posts/{id}
func (d *AdminDeps) EditPost(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Edit post not implemented", http.StatusNotImplemented)
}

// UpdatePost handles PUT /dashboard/posts/{id}
func (d *AdminDeps) UpdatePost(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Update post not implemented", http.StatusNotImplemented)
}

// DeletePost handles DELETE /dashboard/posts/{id}
func (d *AdminDeps) DeletePost(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Delete post not implemented", http.StatusNotImplemented)
}
