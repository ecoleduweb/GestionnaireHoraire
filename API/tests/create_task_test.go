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

func TestCreateActivity(t *testing.T) {

	activity := DTOs.ActivityDTO{
		Name:        "Test tâche",
		Description: "Test automatique de la création de tâche",
		StartDate:   time.Now(),
		EndDate:     time.Now().Add(24 * time.Hour),
		UserId:      doNotDeleteUser.Id,
		ProjectId:   doNotDeleteProject.Id,
		CategoryId:  doNotDeleteCategory.Id,
	}

	w = sendRequest(router, "POST", "/activity", activity)
	assertResponse(t, w, http.StatusCreated, nil)

	// Vérification du corps de la réponse
	var responseBody struct {
		Reponse  string        `json:"reponse"`
		Activity DAOs.Activity `json:"activity"`
	}
	err := json.Unmarshal(w.Body.Bytes(), &responseBody)
	assert.NoError(t, err)
	assert.Equal(t, "La tâche a bien été ajoutée à la base de données.", responseBody.Reponse)
	assert.Equal(t, activity.Name, responseBody.Activity.Name)
	assert.Equal(t, activity.Description, responseBody.Activity.Description)

	// Vérification que la tâche est bien ajoutée dans la base de données
	var createdActivity DAOs.Activity
	errDB := database.DB.Where("name = ?", activity.Name).First(&createdActivity).Error
	assert.NoError(t, errDB)
	assert.Equal(t, activity.Name, createdActivity.Name)
}

func TestDoNotCreateActivityWithEndDateBeforeStartDate(t *testing.T) {

	// Création d'une tâche à envoyer dans la requête
	activity := DTOs.ActivityDTO{
		Name:        "Test tâche",
		Description: "Test automatique de la création de tâche",
		StartDate:   time.Now(),
		EndDate:     time.Now().Add(-24 * time.Hour),
		UserId:      doNotDeleteUser.Id,
		ProjectId:   doNotDeleteProject.Id,
		CategoryId:  doNotDeleteCategory.Id,
	}

	w = sendRequest(router, "POST", "/activity", activity)

	expectedErrors := []DTOs.FieldErrorDTO{
		{Field: "start_date", Message: "La date de début doit être avant la date de fin"},
	}
	assertResponse(t, w, http.StatusBadRequest, expectedErrors)
}

func TestDoNotCreateActivityWithoutNameAndDescription(t *testing.T) {

	activity := DTOs.ActivityDTO{
		Name:        "",
		Description: "",
		StartDate:   time.Now(),
		EndDate:     time.Now().Add(24 * time.Hour),
		UserId:      doNotDeleteUser.Id,
		ProjectId:   doNotDeleteProject.Id,
		CategoryId:  doNotDeleteCategory.Id,
	}

	w = sendRequest(router, "POST", "/activity", activity)

	expectedErrors := []DTOs.FieldErrorDTO{
		{Field: "Name", Message: "Le champ Name est invalide ou manquant"},
		{Field: "Description", Message: "Le champ Description est invalide ou manquant"},
	}
	assertResponse(t, w, http.StatusBadRequest, expectedErrors)
}

func TestDoNotCreateActivityWithLenghtNameOver50(t *testing.T) {

	activity := DTOs.ActivityDTO{
		//Sitation dans le film Astérix et Obélix : mission Cléopâtre
		Name:        "Vous savez, moi je ne crois pas qu’il y ait de bonne ou de mauvaise situation. Moi, si je devais résumer ma vie aujourd’hui avec vous, je dirais que c’est d’abord des rencontres...",
		Description: "Test automatique de la création de tâche",
		StartDate:   time.Now(),
		EndDate:     time.Now().Add(24 * time.Hour),
		UserId:      doNotDeleteUser.Id,
		ProjectId:   doNotDeleteProject.Id,
		CategoryId:  doNotDeleteCategory.Id,
	}
	w = sendRequest(router, "POST", "/activity", activity)

	expectedErrors := []DTOs.FieldErrorDTO{
		{Field: "Name", Message: "Le champ Name est invalide ou manquant"},
	}
	assertResponse(t, w, http.StatusBadRequest, expectedErrors)
}

func TestDoNotCreateActivityWithoutDates(t *testing.T) {

	activity := DTOs.ActivityDTO{
		Name:        "Test tâche",
		Description: "Test automatique de la création de tâche",
		UserId:      doNotDeleteUser.Id,
		ProjectId:   doNotDeleteProject.Id,
		CategoryId:  doNotDeleteCategory.Id,
	}

	w = sendRequest(router, "POST", "/activity", activity)

	expectedErrors := []DTOs.FieldErrorDTO{
		{Field: "StartDate", Message: "Le champ StartDate est invalide ou manquant"},
		{Field: "EndDate", Message: "Le champ EndDate est invalide ou manquant"},
	}
	assertResponse(t, w, http.StatusBadRequest, expectedErrors)
}

func TestDoNotCreateActivityWithInvalidStartDate(t *testing.T) {

	activity := DTOs.ActivityDTO{
		Name:        "Test tâche",
		Description: "Test automatique de la création de tâche",
		StartDate:   time.Time{},
		EndDate:     time.Now().Add(24 * time.Hour),
		UserId:      doNotDeleteUser.Id,
		ProjectId:   doNotDeleteProject.Id,
		CategoryId:  doNotDeleteCategory.Id,
	}

	w = sendRequest(router, "POST", "/activity", activity)

	expectedErrors := []DTOs.FieldErrorDTO{
		{Field: "StartDate", Message: "Le champ StartDate est invalide ou manquant"},
	}
	assertResponse(t, w, http.StatusBadRequest, expectedErrors)
}

func TestDoNotCreateActivityWithInvalidEndDate(t *testing.T) {

	activity := DTOs.ActivityDTO{
		Name:        "Test tâche",
		Description: "Test automatique de la création de tâche",
		StartDate:   time.Now(),
		EndDate:     time.Time{},
		UserId:      doNotDeleteUser.Id,
		ProjectId:   doNotDeleteProject.Id,
		CategoryId:  doNotDeleteCategory.Id,
	}

	w = sendRequest(router, "POST", "/activity", activity)

	expectedErrors := []DTOs.FieldErrorDTO{
		{Field: "EndDate", Message: "Le champ EndDate est invalide ou manquant"},
	}
	assertResponse(t, w, http.StatusBadRequest, expectedErrors)
}
