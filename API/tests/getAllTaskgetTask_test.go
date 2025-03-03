package tests

import (
	"llio-api/controllers"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllTask(t *testing.T) {
	connectDB()
	router := setupTestRouter(http.MethodGet, "/tasks", controllers.GetAllTasks)
	w := sendRequest(router, "GET", "/tasks", nil)

	assertResponse(t, w, http.StatusOK, nil)
	assert.NotNil(t, w.Body)
}

func TestGetTask(t *testing.T) {
	connectDB()
	router := setupTestRouter(http.MethodGet, "/tasks/:id", controllers.GetTaskById)
	w := sendRequest(router, "GET", "/tasks/1", nil)

	assertResponse(t, w, http.StatusOK, nil)
	assert.NotNil(t, w.Body)
}

func TestGetNonExistentTask(t *testing.T) {
	connectDB()
	router := setupTestRouter(http.MethodGet, "/tasks/:id", controllers.GetTaskById)
	w := sendRequest(router, "GET", "/tasks/-1", nil)
	assertResponse(t, w, http.StatusNotFound, nil)
	assert.JSONEq(t, `{"error": "La tâche est introuvable."}`, w.Body.String())
}
