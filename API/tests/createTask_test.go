package tests

import (
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"llio-api/database"
	"llio-api/models/DAOs"
	"llio-api/models/DTOs"

	"github.com/stretchr/testify/assert"
)

// TestCreateTask est la fonction de test pour le contrôleur CreateTask.
func TestCreateTask(t *testing.T) {

	router, w := setupTestRouter()

	// Création d'une tâche à envoyer dans la requête
	task := DTOs.TaskDTO{
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

	w = sendRequest(router, "POST", "/tasks", task)
	assertResponse(t, w, http.StatusCreated, nil)

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

func TestDoNotCreateTaskWithEndDateBeforeStartDate(t *testing.T) {

	router, w := setupTestRouter()

	// Création d'une tâche à envoyer dans la requête
	task := DTOs.TaskDTO{
		Name:        "Test tâche",
		Description: "Test automatique de la création de tâche",
		State:       "à faire",
		Billable:    true,
		StartDate:   time.Now(),
		EndDate:     time.Now().Add(-24 * time.Hour),
		UserId:      1,
		ProjectId:   1,
		CategoryId:  1,
	}

	w = sendRequest(router, "POST", "/tasks", task)

	expectedErrors := []DTOs.FieldErrorDTO{
		{Field: "start_date", Message: "La date de début doit être avant la date de fin"},
	}
	assertResponse(t, w, http.StatusBadRequest, expectedErrors)
}

func TestDoNotCreateTaskWithoutNameAndDescription(t *testing.T) {

	router, w := setupTestRouter()

	// Création d'une tâche à envoyer dans la requête
	task := DTOs.TaskDTO{
		Name:        "",
		Description: "",
		State:       "à faire",
		Billable:    true,
		StartDate:   time.Now(),
		EndDate:     time.Now().Add(24 * time.Hour),
		UserId:      1,
		ProjectId:   1,
		CategoryId:  1,
	}

	w = sendRequest(router, "POST", "/tasks", task)

	expectedErrors := []DTOs.FieldErrorDTO{
		{Field: "Name", Message: "Le champ Name est invalide ou manquant"},
		{Field: "Description", Message: "Le champ Description est invalide ou manquant"},
	}
	assertResponse(t, w, http.StatusBadRequest, expectedErrors)
}

func TestDoNotCreateTaskWithLenghtNameOver50(t *testing.T) {

	router, w := setupTestRouter()

	// Création d'une tâche à envoyer dans la requête
	task := DTOs.TaskDTO{
		//Sitation dans le film Astérix et Obélix : mission Cléopâtre
		Name:        "Vous savez, moi je ne crois pas qu’il y ait de bonne ou de mauvaise situation. Moi, si je devais résumer ma vie aujourd’hui avec vous, je dirais que c’est d’abord des rencontres...",
		Description: "Test automatique de la création de tâche",
		State:       "à faire",
		Billable:    true,
		StartDate:   time.Now(),
		EndDate:     time.Now().Add(24 * time.Hour),
		UserId:      1,
		ProjectId:   1,
		CategoryId:  1,
	}
	w = sendRequest(router, "POST", "/tasks", task)

	expectedErrors := []DTOs.FieldErrorDTO{
		{Field: "Name", Message: "Le champ Name est invalide ou manquant"},
	}
	assertResponse(t, w, http.StatusBadRequest, expectedErrors)
}

func TestDoNotCreateTaskWithoutDates(t *testing.T) {

	router, w := setupTestRouter()

	// Création d'une tâche à envoyer dans la requête
	task := DTOs.TaskDTO{
		Name:        "Test tâche",
		Description: "Test automatique de la création de tâche",
		State:       "à faire",
		Billable:    true,
		UserId:      1,
		ProjectId:   1,
		CategoryId:  1,
	}

	w = sendRequest(router, "POST", "/tasks", task)

	expectedErrors := []DTOs.FieldErrorDTO{
		{Field: "StartDate", Message: "Le champ StartDate est invalide ou manquant"},
		{Field: "EndDate", Message: "Le champ EndDate est invalide ou manquant"},
	}
	assertResponse(t, w, http.StatusBadRequest, expectedErrors)
}

func TestDoNotCreateTaskWithInvalidStartDate(t *testing.T) {

	router, w := setupTestRouter()

	// Création d'une tâche à envoyer dans la requête
	task := DTOs.TaskDTO{
		Name:        "Test tâche",
		Description: "Test automatique de la création de tâche",
		State:       "à faire",
		Billable:    true,
		StartDate:   time.Time{},
		EndDate:     time.Now().Add(24 * time.Hour),
		UserId:      1,
		ProjectId:   1,
		CategoryId:  1,
	}

	w = sendRequest(router, "POST", "/tasks", task)

	expectedErrors := []DTOs.FieldErrorDTO{
		{Field: "StartDate", Message: "Le champ StartDate est invalide ou manquant"},
	}
	assertResponse(t, w, http.StatusBadRequest, expectedErrors)
}

func TestDoNotCreateTaskWithInvalidEndDate(t *testing.T) {

	router, w := setupTestRouter()

	// Création d'une tâche à envoyer dans la requête
	task := DTOs.TaskDTO{
		Name:        "Test tâche",
		Description: "Test automatique de la création de tâche",
		State:       "à faire",
		Billable:    true,
		StartDate:   time.Now(),
		EndDate:     time.Time{},
		UserId:      1,
		ProjectId:   1,
		CategoryId:  1,
	}

	w = sendRequest(router, "POST", "/tasks", task)

	expectedErrors := []DTOs.FieldErrorDTO{
		{Field: "EndDate", Message: "Le champ EndDate est invalide ou manquant"},
	}
	assertResponse(t, w, http.StatusBadRequest, expectedErrors)
}
