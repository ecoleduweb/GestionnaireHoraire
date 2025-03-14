package tests

import (
	"llio-api/controllers"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllActivity(t *testing.T) {
	connectDB()
	router := setupTestRouter(http.MethodGet, "/activities", controllers.GetAllActivities)
	w := sendRequest(router, "GET", "/activities", nil)

	assertResponse(t, w, http.StatusOK, nil)
	assert.NotNil(t, w.Body)
}

func TestGetActivity(t *testing.T) {
	connectDB()
	router := setupTestRouter(http.MethodGet, "/activities/:id", controllers.GetActivityById)
	w := sendRequest(router, "GET", "/activities/1", nil)

	assertResponse(t, w, http.StatusOK, nil)
	assert.NotNil(t, w.Body)
}

func TestGetNonExistentActivity(t *testing.T) {
	connectDB()
	router := setupTestRouter(http.MethodGet, "/activities/:id", controllers.GetActivityById)
	w := sendRequest(router, "GET", "/activities/-1", nil)
	assertResponse(t, w, http.StatusNotFound, nil)
	assert.JSONEq(t, `{"error": "La tâche est introuvable."}`, w.Body.String())
}
