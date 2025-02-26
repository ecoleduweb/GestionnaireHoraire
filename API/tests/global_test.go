package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"llio-api/database"
	"llio-api/models/DTOs"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
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

	// Vérifier si le répertoire courant contient déjà le dossier "API"
	if !strings.HasSuffix(currentDir, "API") {
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
	} else {
		fmt.Println("Déjà dans le répertoire API.")
	}
}

// setupTestRouter initialise un routeur de test et un enregistreur de réponse
func setupTestRouter(method, route string, controller gin.HandlerFunc) (*gin.Engine, *httptest.ResponseRecorder) {
	router := gin.Default()

	switch method {
	case http.MethodGet:
		router.GET(route, controller)
	case http.MethodPost:
		router.POST(route, controller)
	case http.MethodPut:
		router.PUT(route, controller)
	case http.MethodDelete:
		router.DELETE(route, controller)
	// Ajoutez d'autres méthodes HTTP si nécessaire
	default:
		panic("Méthode HTTP non supportée")
	}

	return router, httptest.NewRecorder()
}

// connectDB se connecte à la base de données
func connectDB() {
	changeCurrentDiretory()
	os.Setenv("ENV", "TEST")
	database.Connect()
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
