package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"llio-api/models/enums"

	"github.com/stretchr/testify/assert"
)

func TestGetAllManagers(t *testing.T) {
	w := sendRequest(router, "GET", "/users/managers", nil, enums.ProjectManager)
	assertResponse(t, w, http.StatusOK, nil)
	assert.NotNil(t, w.Body)
}

func TestGetAllManagersWithInvalidToken(t *testing.T) {
	// Create a request without using the sendRequest helper to avoid token generation
	req, _ := http.NewRequest("GET", "/users/managers", nil)
	req.Header.Set("Content-Type", "application/json")

	// Add an invalid token cookie
	cookie := &http.Cookie{
		Name:     "access_token",
		Value:    "invalid_token",
		HttpOnly: true,
		Path:     "/",
	}
	req.AddCookie(cookie)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assertResponse(t, w, http.StatusUnauthorized, nil)
}

func TestGetAllManagersWithoutPermission(t *testing.T) {
	w := sendRequest(router, "GET", "/users/managers", nil, enums.Employee)
	assertResponse(t, w, http.StatusForbidden, nil)
}
