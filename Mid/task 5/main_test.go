package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBookingValidation(t *testing.T) {
	router := SetupRouter()

	tests := []struct {
		name           string
		payload        map[string]interface{}
		expectedStatus int
	}{
		{
			name: "Valid request",
			payload: map[string]interface{}{
				"user_id":  1,
				"table_id": 5,
				"email":    "test@example.com",
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "Missing email",
			payload: map[string]interface{}{
				"user_id":  1,
				"table_id": 5,
			},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "Invalid table ID (zero)",
			payload: map[string]interface{}{
				"user_id":  1,
				"table_id": 0,
				"email":    "test@example.com",
			},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "Bad email format",
			payload: map[string]interface{}{
				"user_id":  1,
				"table_id": 2,
				"email":    "not-an-email",
			},
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsonPayload, _ := json.Marshal(tt.payload)
			req, _ := http.NewRequest("POST", "/book", bytes.NewBuffer(jsonPayload))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)
		})
	}
}
