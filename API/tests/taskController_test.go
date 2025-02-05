package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"llio-api/controllers"
	"llio-api/database"
	"llio-api/models"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// TestCreateTask est la fonction de test pour le contrôleur CreateTask.
func TestCreateTask(t *testing.T) {

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

	// Initialiser la base de données de test
	database.Connect()

	// Création d'un routeur Gin
	router := gin.Default()
	router.POST("/tasks/create", controllers.CreateTask)

	// Création d'une tâche à envoyer dans la requête
	task := struct {
		Name        string    `json:"name"`
		Description string    `json:"description"`
		State       string    `json:"state"`
		Billable    bool      `json:"billable"`
		StartDate   time.Time `json:"start_date"`
		EndDate     time.Time `json:"end_date"`
		UserId      int       `json:"user_id"`
		ProjectId   int       `json:"project_id"`
		CategoryId  int       `json:"category_id"`
	}{
		Name:        "Test Task",
		Description: "This is a test task",
		State:       "à faire",
		Billable:    true,
		StartDate:   time.Now(),
		EndDate:     time.Now().Add(24 * time.Hour),
		UserId:      1,
		ProjectId:   1,
		CategoryId:  1,
	}

	// Conversion de la tâche en JSON
	jsonValue, _ := json.Marshal(task)

	// Création d'une requête HTTP POST
	req, _ := http.NewRequest("POST", "/tasks/create", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	// Création d'un enregistreur de réponse HTTP
	w := httptest.NewRecorder()

	// Exécution de la requête
	router.ServeHTTP(w, req)

	// Vérification du code de statut HTTP
	assert.Equal(t, http.StatusOK, w.Code)

	// Vérification du corps de la réponse
	expectedResponse := `{"reponse":"La tâche a bien été ajoutée à la base de données."}`
	assert.Equal(t, expectedResponse, w.Body.String())

	// Vérification que la tâche est bien ajoutée dans la base de données
	var createdTask models.Task
	errDB := database.DB.Where("name = ?", "Test Task").First(&createdTask).Error
	assert.NoError(t, errDB)
	assert.Equal(t, "Test Task", createdTask.Name)
}
