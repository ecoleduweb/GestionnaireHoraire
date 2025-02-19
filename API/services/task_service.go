package services

import (
	"llio-api/models/DAOs"
	"llio-api/models/DTOs"
	"llio-api/repositories"
)

func VerifyCreateTaskJSON(taskDTO *DTOs.TaskDTO) []DTOs.FieldError {
	var errors []DTOs.FieldError

	if taskDTO.State == "" {
		taskDTO.State = "à faire"
	}

	// Vérifier que StartDate est avant EndDate
	if taskDTO.StartDate.After(taskDTO.EndDate) {
		errors = append(errors, DTOs.FieldError{
			Field:   "start_date",
			Message: "La date de début doit être avant la date de fin",
		})
	}

	return errors
}

// CreateTaskService encapsule la logique métier pour la création d'une tâche
func CreateTaskService(taskDTO *DTOs.TaskDTO) (*DTOs.TaskResponse, error) {

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

	err := repositories.CreateTask(task)
	if err != nil {
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
