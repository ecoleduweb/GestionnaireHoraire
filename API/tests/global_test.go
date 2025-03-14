package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"llio-api/controllers"
	"llio-api/database"
	"llio-api/models/DAOs"
	"llio-api/models/DTOs"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// variables globales utilisées pour led tests
var (
	router              *gin.Engine
	w                   *httptest.ResponseRecorder
	doNotDeleteUser     DAOs.User
	doNotDeleteCategory DAOs.Project
	doNotDeleteProject  DAOs.Category
)

// agit comme un gros before each
func TestMain(m *testing.M) {
	r, recorder := setupTestRouter()
	router = r
	w = recorder

	//Ajoute des enregistrements de tests
	prepareTestData()
	// Permet de rouler tous les tests. Aucun test ne fonctionne sans  ça.
	exitCode := m.Run()

	// TODO vider la BD après les tests.
	drop_all_tables()
	// affiche le code de sortie des tests
	os.Exit(exitCode)
}

func prepareTestData() {
	testUser := DAOs.User{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@example.com",
	}
	database.DB.Create(&testUser)
	doNotDeleteUser = testUser
	testProject := DAOs.Project{
		Name:        "Test Project",
		Description: "Sample project",
		Status:      true,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		EndAt:       time.Now(),
	}
	database.DB.Create(&testProject)
	doNotDeleteCategory = testProject

	testCategory := DAOs.Category{
		Name:        "Test Category",
		Description: "Sample category",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	database.DB.Create(&testCategory)
	doNotDeleteProject = testCategory
}

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
func setupTestRouter() (*gin.Engine, *httptest.ResponseRecorder) {
	changeCurrentDiretory()
	os.Setenv("ENV", "TEST")
	database.Connect()

	router := gin.Default()
	router.POST("/activity", controllers.CreateActivity)

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

func drop_all_tables() {
	// Supprimer toutes les tables de façon vraiment dégueulasse
	database.DB.Exec("SET FOREIGN_KEY_CHECKS = 0")
	database.DB.Exec("DROP TABLE IF EXISTS schema_migrations")
	database.DB.Exec("DROP TABLE IF EXISTS users")
	database.DB.Exec("DROP TABLE IF EXISTS projects")
	database.DB.Exec("DROP TABLE IF EXISTS categories")
	database.DB.Exec("DROP TABLE IF EXISTS activities")
	database.DB.Exec("SET FOREIGN_KEY_CHECKS = 1")
}
