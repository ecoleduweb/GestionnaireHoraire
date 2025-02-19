package services

import (
	"errors"
	"llio-api/models/DAOs"
	"llio-api/models/DTOs"
	"llio-api/repositories"
)

// CreateTaskService encapsule la logique métier pour la création d'une tâche
func CreateTaskService(taskDTO *DTOs.TaskDTO) (*DTOs.TaskResponse, error) {
	//Vérification des champs obligatoire
	if taskDTO.Name == "" || taskDTO.Description == "" || taskDTO.StartDate.IsZero() || taskDTO.EndDate.IsZero() {
		return nil, errors.New("des champs obligatoires sont manquants")
	}

	// Mapper le body vers le modèle Task
	task := &DAOs.Task{
		Name:        taskDTO.Name,
		Description: taskDTO.Description,
		State:       taskDTO.State,
		Billable:    taskDTO.Billable,
		StartDate:   taskDTO.StartDate,
		EndDate:     taskDTO.EndDate,
		UserId:      taskDTO.UserId,
		ProjectId:   taskDTO.ProjectId,
		CategoryId:  taskDTO.CategoryId,
	}

	if task.State == "" {
		task.State = "à faire"
	}

	if err := repositories.CreateTask(task); err != nil {
		return nil, err
	}

	taskResponse := &DTOs.TaskResponse{
		Name:        task.Name,
		Description: task.Description,
		State:       task.State,
		Billable:    task.Billable,
		StartDate:   task.StartDate,
		EndDate:     task.EndDate,
	}
	return taskResponse, nil
}
