package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

func generateTestToken(secret []byte, expired bool) string {
	expiration := time.Now().Add(time.Hour)
	if expired {
		expiration = time.Now().Add(-time.Hour)
	}

	claims := jwt.MapClaims{
		"sub": "test_user",
		"exp": expiration.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString(secret)
	return tokenString
}

func TestJWTAuthMiddleware(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Success with valid token", func(t *testing.T) {
		r := gin.New()
		r.GET("/test", JWTAuthMiddleware(), func(c *gin.Context) {
			c.Status(http.StatusOK)
		})

		token := generateTestToken(secretKey, false)
		req, _ := http.NewRequest("GET", "/test", nil)
		req.Header.Set("Authorization", "Bearer "+token)

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("Fail with expired token", func(t *testing.T) {
		r := gin.New()
		r.GET("/test", JWTAuthMiddleware(), func(c *gin.Context) {
			c.Status(http.StatusOK)
		})

		token := generateTestToken(secretKey, true)
		req, _ := http.NewRequest("GET", "/test", nil)
		req.Header.Set("Authorization", "Bearer "+token)

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})

	t.Run("Fail without header", func(t *testing.T) {
		r := gin.New()
		r.GET("/test", JWTAuthMiddleware(), func(c *gin.Context) {
			c.Status(http.StatusOK)
		})

		req, _ := http.NewRequest("GET", "/test", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})
}
