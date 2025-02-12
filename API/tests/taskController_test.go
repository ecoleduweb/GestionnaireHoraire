package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"llio-api/controllers"
	"llio-api/database"
	"llio-api/models/DAOs"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// TestCreateTask est la fonction de test pour le contrôleur CreateTask.
func TestCreateTask(t *testing.T) {

	changeCurrentDiretory()

	os.Setenv("ENV", "DEV")
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
		Name:        "Test tâche",
		Description: "Test automatique de la création de tâche",
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
	assert.Equal(t, http.StatusCreated, w.Code)

	// Vérification du corps de la réponse
	var responseBody struct {
		Reponse string    `json:"reponse"`
		Task    DAOs.Task `json:"task"`
	}
	err := json.Unmarshal(w.Body.Bytes(), &responseBody)
	assert.NoError(t, err)
	assert.Equal(t, "La tâche a bien été ajoutée à la base de données.", responseBody.Reponse)
	assert.Equal(t, task.Name, responseBody.Task.Name)
	assert.Equal(t, task.Description, responseBody.Task.Description)
	assert.Equal(t, task.State, responseBody.Task.State)
	assert.Equal(t, task.Billable, responseBody.Task.Billable)

	// Vérification que la tâche est bien ajoutée dans la base de données
	var createdTask DAOs.Task
	errDB := database.DB.Where("name = ?", task.Name).First(&createdTask).Error
	assert.NoError(t, errDB)
	assert.Equal(t, task.Name, createdTask.Name)
}
