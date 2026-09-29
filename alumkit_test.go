package alumkit

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAppServeHTTP(t *testing.T) {
	t.Parallel()

	app := &App{router: newRouter()}

	called := false
	app.Get("/test", func(w http.ResponseWriter, r *http.Request) {
		called = true
		w.WriteHeader(http.StatusOK)
	})

	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	w := httptest.NewRecorder()

	app.ServeHTTP(w, req)

	if !called {
		t.Error("handler was not called")
	}
	if w.Code != http.StatusOK {
		t.Errorf("status = %d, want %d", w.Code, http.StatusOK)
	}
}

func TestAppHTTPMethods(t *testing.T) {
	t.Parallel()

	app := &App{router: newRouter()}

	methods := []struct {
		method string
		path   string
	}{
		{http.MethodGet, "/get"},
		{http.MethodPost, "/post"},
		{http.MethodPut, "/put"},
		{http.MethodDelete, "/delete"},
	}

	// Register all routes first
	for _, m := range methods {
		app.router.Method(m.method, m.path, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(r.Method))
		}))
	}

	// Then make requests
	for _, m := range methods {
		t.Run(m.method, func(t *testing.T) {
			t.Parallel()

			req := httptest.NewRequest(m.method, m.path, nil)
			w := httptest.NewRecorder()
			app.ServeHTTP(w, req)

			if w.Body.String() != m.method {
				t.Errorf("body = %q, want %q", w.Body.String(), m.method)
			}
		})
	}
}
