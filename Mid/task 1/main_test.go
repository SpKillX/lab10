package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEndpoints(t *testing.T) {
	router := SetupRouter()

	tests := []struct {
		name           string
		url            string
		expectedStatus int
	}{
		{"Test Ping", "/ping", http.StatusOK},
		{"Test Status", "/status", http.StatusOK},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", tt.url, nil)
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)
		})
	}
}
