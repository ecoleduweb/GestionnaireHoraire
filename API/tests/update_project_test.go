package tests

import (
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"llio-api/models/DAOs"
	"llio-api/models/DTOs"
	"llio-api/models/enums"

	"github.com/stretchr/testify/assert"
)

func TestUpdateProject(t *testing.T) {
	updatedProject := DTOs.ProjectDTO{
		Id:          doNotDeleteProject2.Id,
		ManagerId:   doNotDeleteUser.Id,
		Name:        "Projet mis à jour",
		Description: "Description mise à jour du projet",
		Status:      enums.ProjectStatus(enums.InProgress),
		// On laisse les autres champs tels quels
	}

	w := sendRequest(router, "PUT", "/project", updatedProject, enums.Administrator)
	assertResponse(t, w, http.StatusOK, nil)

	var responseBody struct {
		UpdatedProject DAOs.Project `json:"updatedProject"`
	}
	err := json.Unmarshal(w.Body.Bytes(), &responseBody)
	assert.NoError(t, err)
	assert.Equal(t, doNotDeleteProject2.Id, responseBody.UpdatedProject.Id)
	assert.Equal(t, updatedProject.Name, responseBody.UpdatedProject.Name)
	assert.Equal(t, updatedProject.Description, responseBody.UpdatedProject.Description)
}

func TestUpdateProjectStatus(t *testing.T) {
	updatedProject := DTOs.ProjectDTO{
		Id:          doNotDeleteProject2.Id,
		ManagerId:   doNotDeleteUser.Id,
		Name:        doNotDeleteProject2.Name,
		Description: doNotDeleteProject2.Description,
		Status:      enums.ProjectStatus(enums.Finish),
	}

	w := sendRequest(router, "PUT", "/project", updatedProject, enums.Administrator)
	assertResponse(t, w, http.StatusOK, nil)

	var responseBody struct {
		UpdatedProject DAOs.Project `json:"updatedProject"`
	}
	err := json.Unmarshal(w.Body.Bytes(), &responseBody)
	assert.NoError(t, err)
	assert.Equal(t, enums.ProjectStatus(enums.Finish), responseBody.UpdatedProject.Status)
}

func TestUpdateProjectEndDate(t *testing.T) {
	// Date de fin future
	endDate := time.Now().Add(30 * 24 * time.Hour) // 30 jours dans le futur

	// Modification de la date de fin du projet
	updatedProject := DTOs.ProjectDTO{
		Id:          doNotDeleteProject2.Id,
		ManagerId:   doNotDeleteUser.Id,
		Name:        doNotDeleteProject2.Name,
		Description: doNotDeleteProject2.Description,
		Status:      doNotDeleteProject2.Status,
		EndAt:       endDate,
	}

	w := sendRequest(router, "PUT", "/project", updatedProject, enums.Administrator)
	assertResponse(t, w, http.StatusOK, nil)

	var responseBody struct {
		UpdatedProject DAOs.Project `json:"updatedProject"`
	}
	err := json.Unmarshal(w.Body.Bytes(), &responseBody)
	assert.NoError(t, err)
	assert.NotNil(t, responseBody.UpdatedProject.EndAt)
}

func TestDoNotUpdateProjectWithNonExistingId(t *testing.T) {
	nonExistingProject := DTOs.ProjectDTO{
		Id:          999999, // ID qui n'existe probablement pas
		ManagerId:   doNotDeleteUser.Id,
		Name:        "Projet inexistant",
		Description: "Description d'un projet inexistant",
		Status:      enums.ProjectStatus(enums.InProgress),
	}

	w := sendRequest(router, "PUT", "/project", nonExistingProject, enums.Administrator)
	assertResponse(t, w, http.StatusNotFound, nil)
}

func TestDoNotUpdateProjectWithInvalidName(t *testing.T) {

	// Modification du projet avec un nom invalide
	updatedProject := DTOs.ProjectDTO{
		Id:          doNotDeleteProject2.Id,
		ManagerId:   doNotDeleteUser.Id,
		Name:        "", // Nom vide
		Description: doNotDeleteProject2.Description,
		Status:      doNotDeleteProject2.Status,
	}

	w := sendRequest(router, "PUT", "/project", updatedProject, enums.Administrator)
	expectedErrors := []DTOs.FieldErrorDTO{
		{Field: "name", Message: "Le champ name est invalide ou manquant"},
	}
	assertResponse(t, w, http.StatusBadRequest, expectedErrors)
}

func TestDoNotUpdateProjectWithInconsistentDates(t *testing.T) {
	// Temps actuel
	now := time.Now()

	// Modification du projet avec des dates incohérentes
	updatedProject := DTOs.ProjectDTO{
		Id:          doNotDeleteProject2.Id,
		ManagerId:   doNotDeleteUser.Id,
		Name:        doNotDeleteProject2.Name,
		Description: doNotDeleteProject2.Description,
		Status:      doNotDeleteProject2.Status,
		CreatedAt:   now,
		EndAt:       now.Add(-24 * time.Hour), // Un jour avant CreatedAt
	}

	w := sendRequest(router, "PUT", "/project", updatedProject, enums.Administrator)
	expectedErrors := []DTOs.FieldErrorDTO{
		{Field: "endAt", Message: "La date de fin doit être après la date de création"},
	}
	assertResponse(t, w, http.StatusBadRequest, expectedErrors)
}
