package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHomeHandler_OK(t *testing.T) {
	// mock request for server, passed to handler as request object
	req := httptest.NewRequest("GET", "/", nil)
	// responsewriter interface and records all responses from the handler
	w := httptest.NewRecorder()
	// variables to call  HomeHandler function
	HomeHandler(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}
}
