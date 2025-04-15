package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"llio-api/database"
	"llio-api/handlers"
	"llio-api/models/DAOs"
	"llio-api/models/DTOs"
	"llio-api/models/enums"
	"llio-api/routes"
	"llio-api/services"
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
	doNotDeleteCategory DAOs.Category
	doNotDeleteProject  DAOs.Project
	doNotDeleteProject2 DAOs.Project
	doNotDeleteActivity DAOs.Activity
)

// Global JWT token for authentication in tests
var accessToken string

// Create and set JWT token for tests

func createAndSetAccessToken(role enums.UserRole) {
	// Create a JWT token for the test user
	token, err := services.CreateJWTToken(doNotDeleteUser.Email, doNotDeleteUser.FirstName, doNotDeleteUser.LastName, time.Now().Add(time.Hour), role)
	if err != nil {
		log.Fatalf("Failed to create JWT token: %v", err)
	}
	accessToken = token
}

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
		Status:      enums.ProjectStatus(enums.InProgress),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		EndAt:       time.Now(),
	}
	database.DB.Create(&testProject)
	doNotDeleteProject = testProject
	testProject2 := DAOs.Project{
		Name:        "Test Project 2",
		Description: "Sample project 2",
		Status:      enums.ProjectStatus(enums.InProgress),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		EndAt:       time.Now(),
	}
	database.DB.Create(&testProject2)
	doNotDeleteProject2 = testProject2

	testCategory := DAOs.Category{
		Name:        "Test Category",
		Description: "Sample category",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		UserId:      doNotDeleteUser.Id,
		ProjectId:   doNotDeleteProject.Id,
	}
	database.DB.Create(&testCategory)
	doNotDeleteCategory = testCategory

	testActivity := DAOs.Activity{
		Name:        "Test Activity",
		StartDate:   time.Now(),
		EndDate:     time.Now().Add(time.Hour),
		Description: "test description",
		UserId:      doNotDeleteUser.Id,
		ProjectId:   doNotDeleteProject.Id,
		CategoryId:  doNotDeleteCategory.Id,
	}
	database.DB.Create(&testActivity)
	doNotDeleteActivity = testActivity
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
	// used for the api routes
	routes.RegisterRoutes(router)
	// used for the health check
	handlers.ApiStatus(router)

	return router, httptest.NewRecorder()
}

// sendRequest envoie une requête HTTP au routeur de test
// pour créer une requête http avec un role administrateur, on ajoute le role voulu à la fin : sendRequest(router, "POST", "/activity", activity, enums.Employee)
func sendRequest(router *gin.Engine, method, url string, body interface{}, userRole ...enums.UserRole) *httptest.ResponseRecorder {
	var req *http.Request
	// If accessToken exists, we need to add it to the request cookies
	// This will be used in non-authenticated helper functions
	if method == "GET" || body == nil {
		// Pour GET ou body nil, ne pas inclure de corps
		req, _ = http.NewRequest(method, url, nil)
	} else {
		// Pour les autres méthodes avec body non-nil
		jsonValue, _ := json.Marshal(body)
		req, _ = http.NewRequest(method, url, bytes.NewBuffer(jsonValue))
	}

	if (userRole != nil) && len(userRole) > 0 {
		createAndSetAccessToken(userRole[0])
	} else {
		createAndSetAccessToken(enums.Employee)
	}

	cookie := &http.Cookie{
		Name:     "access_token",
		Value:    accessToken,
		HttpOnly: true,
		Path:     "/",
	}
	// Add cookies to the request if any exist
	req.AddCookie(cookie)
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
