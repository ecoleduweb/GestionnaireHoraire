package services

import (
	"llio-api/models/DAOs"
	"llio-api/models/DTOs"
	"llio-api/repositories"
)

func VerifyCreateTaskJSON(taskDTO *DTOs.TaskDTO) []DTOs.FieldErrorDTO {
	var errors []DTOs.FieldErrorDTO

	// Vérifier que StartDate est avant EndDate
	if taskDTO.StartDate.After(taskDTO.EndDate) {
		errors = append(errors, DTOs.FieldErrorDTO{
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
		Billable:    task.Billable,
		StartDate:   task.StartDate,
		EndDate:     task.EndDate,
	}
	return taskResponse, nil
}

// Fonction qui retourne toutes les tâches
// Convertion des tâches DAOs en tâches DTOs
func GetAllTasksService() ([]*DTOs.TaskDTO, error) {
	tasks, err := repositories.GetAllTasks()
	if err != nil {
		return nil, err
	}

	var tasksDTOs []*DTOs.TaskDTO
	for _, task := range tasks {
		taskDTO := &DTOs.TaskDTO{
			Id:          task.Id,
			Name:        task.Name,
			Description: task.Description,
			Billable:    task.Billable,
			StartDate:   task.StartDate,
			EndDate:     task.EndDate,
			UserId:      task.UserId,
			ProjectId:   task.ProjectId,
			CategoryId:  task.CategoryId,
		}
		tasksDTOs = append(tasksDTOs, taskDTO)
	}

	return tasksDTOs, nil
}

// Fonction qui retourne une tâche par son id
// Convertion de la tâche DAOs en tâche DTOs
func GetTaskByIdService(id string) (*DTOs.TaskDTO, error) {
	task, err := repositories.GetTaskById(id)
	if err != nil {
		return nil, err
	} else if task == nil {
		return nil, nil
	}

	taskDTO := &DTOs.TaskDTO{
		Id:          task.Id,
		Name:        task.Name,
		Description: task.Description,
		Billable:    task.Billable,
		StartDate:   task.StartDate,
		EndDate:     task.EndDate,
		UserId:      task.UserId,
		ProjectId:   task.ProjectId,
		CategoryId:  task.CategoryId,
	}

	return taskDTO, nil
}

func UpdateTaskService(taskDTO *DTOs.TaskResponse) error {

	taskDAO := &DAOs.Task{
		Id:          taskDTO.Id,
		Name:        taskDTO.Name,
		Description: taskDTO.Description,
		Billable:    taskDTO.Billable,
		StartDate:   taskDTO.StartDate,
		EndDate:     taskDTO.EndDate,
	}
	return repositories.UpdateTask(taskDAO)
}
