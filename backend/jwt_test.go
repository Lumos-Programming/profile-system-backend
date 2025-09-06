package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestJWT(t *testing.T) {
	token := generateDummyJWT("test_user")
	if token == "" {
		t.Errorf("Failed to generate JWT token")
		return
	}

	req, _ := http.NewRequest("GET", "/api/me", nil)
	req.AddCookie(&http.Cookie{
		Name:  authCookieName,
		Value: token,
	})

	router := gin.Default()
	router.Use(authHandler)
	router.GET("/api/me", func(c *gin.Context) {
		userID := c.GetString("user_id")
		c.JSON(200, gin.H{"user_id": userID})
	})

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	data := w.Body.String()
	if w.Code != 200 || data != `{"user_id":"test_user"}` {
		t.Errorf("Unexpected response: %d %s", w.Code, data)
	}
}
