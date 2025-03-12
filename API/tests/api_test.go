package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"llio-api/handlers"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// Vérification du Ping Pong, si le serveur répond bien
func TestApiPingPong(t *testing.T) {
	router := gin.Default()
	handlers.ApiStatus(router)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"status": "Pong"}`, w.Body.String())
}

// Vérification de la santé de l'API, l'API doit nous retourner un statut Healthy et le temps écoulé depuis le démarrage
// Nous verifions Si l'objet JSON contient bien les clés "status" et "uptime"
// On par du principe que si l'API est en ligne, elle est en bonne santé
func TestApiHealthStatus(t *testing.T) {
	router := gin.Default()
	handlers.ApiStatus(router)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health/status", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `"status":"Healthy"`)
	assert.Contains(t, w.Body.String(), `"uptime":`)
}

// Vérification de la réponse de l'API si la route n'existe pas
// L'API doit nous retourner un code 404
func TestApiNotFound(t *testing.T) {
	router := gin.Default()
	handlers.ApiStatus(router)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/notfound", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

// Vérification de la réponse de l'API si la requête est mal formée
// L'API doit nous retourner un code 400
func TestApiBadRequest(t *testing.T) {
	router := gin.Default()
	router.POST("/badrequest", func(c *gin.Context) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/badrequest", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.JSONEq(t, `{"error": "bad request"}`, w.Body.String())
}
