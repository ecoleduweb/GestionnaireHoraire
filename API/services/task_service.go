package services

import (
	"errors"
	"llio-api/models/DAOs"
	"llio-api/models/DTOs"
	"llio-api/repositories"
)

// CreateTaskService encapsule la logique métier pour la création d'une tâche
func CreateTaskService(reqBody *DTOs.TaskDTO) (*DAOs.Task, error) {
	//Vérification des champs obligatoire
	if reqBody.Name == "" || reqBody.Description == "" || reqBody.StartDate.IsZero() || reqBody.EndDate.IsZero() {
		return nil, errors.New("des champs obligatoires sont manquants")
	}

	// Mapper le body vers le modèle Task
	task := &DAOs.Task{
		Name:        reqBody.Name,
		Description: reqBody.Description,
		State:       reqBody.State,
		Billable:    reqBody.Billable,
		StartDate:   reqBody.StartDate,
		EndDate:     reqBody.EndDate,
		UserId:      reqBody.UserId,
		ProjectId:   reqBody.ProjectId,
		CategoryId:  reqBody.CategoryId,
	}

	// Appliquer une valeur par défaut si nécessaire
	if task.State == "" {
		task.State = "à faire"
	}

	// Appeler la couche repository pour insérer la tâche
	if err := repositories.CreateTask(task); err != nil {
		return nil, err
	}

	return task, nil
}
