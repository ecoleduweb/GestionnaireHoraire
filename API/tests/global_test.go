package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"llio-api/controllers"
	"llio-api/database"
	"llio-api/models/DTOs"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// Change de répertoir pour trouver le .env
func changeCurrentDiretory() {
	// Obtenir le répertoire courant actuel
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Erreur lors de la récupération du répertoire courant :", err)
		return
	}
	fmt.Println("Répertoire courant actuel :", currentDir)

	// Remonter d'un niveau dans l'arborescence des répertoires pour trouver le .env
	err = os.Chdir("..")
	if err != nil {
		log.Fatalf("Erreur lors du changement de répertoire :%v", err)
		return
	}

	// Vérifier le nouveau répertoire courant
	updatedDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Erreur lors de la récupération du répertoire courant :", err)
		return
	}
	fmt.Println("Nouveau répertoire courant :", updatedDir)
}

// setupTestRouter initialise un routeur de test et un enregistreur de réponse
func setupTestRouter() (*gin.Engine, *httptest.ResponseRecorder) {
	changeCurrentDiretory()
	os.Setenv("ENV", "DEV")
	database.Connect()

	router := gin.Default()
	router.POST("/tasks", controllers.CreateTask)

	return router, httptest.NewRecorder()
}

// sendRequest envoie une requête HTTP au routeur de test
func sendRequest(router *gin.Engine, method, url string, body interface{}) *httptest.ResponseRecorder {
	jsonValue, _ := json.Marshal(body)
	req, _ := http.NewRequest(method, url, bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	return w
}

// assertResponse vérifie si le code de statut de la réponse et les erreurs correspondent à ceux attendus
func assertResponse(t *testing.T, w *httptest.ResponseRecorder, expectedStatus int, expectedErrors []DTOs.FieldErrorDTO) {
	assert.Equal(t, expectedStatus, w.Code)

	if expectedErrors != nil {
		var responseBody struct {
			Errors []DTOs.FieldErrorDTO `json:"errors"`
		}
		err := json.Unmarshal(w.Body.Bytes(), &responseBody)
		assert.NoError(t, err)
		assert.Equal(t, expectedErrors, responseBody.Errors)
	}
}
