package services

import (
	"llio-api/models/DAOs"
	"llio-api/models/DTOs"
	"llio-api/repositories"
	"log"

	"github.com/jinzhu/copier"
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

// CreateTask encapsule la logique métier pour la création d'une tâche
func CreateTask(taskDTO *DTOs.TaskDTO) (*DTOs.TaskDTO, error) {

	// Mapper le body vers le modèle Task

	task := &DAOs.Task{}
	err := copier.Copy(task, taskDTO)
	if err != nil {
		log.Printf("Erreur lors de la copie des champs: %v", err)
		return nil, err
	}

	taskDAOAded, err := repositories.CreateTask(task)
	if err != nil {
		return nil, err
	}

	taskDTOResponse := &DTOs.TaskDTO{}
	err = copier.Copy(taskDTOResponse, taskDAOAded)
	if err != nil {
		log.Printf("Erreur lors de la copie des champs: %v", err)
		return nil, err
	}
	return taskDTOResponse, nil
}

// Fonction qui retourne toutes les tâches
// Convertion des tâches DAOs en tâches DTOs
func GetAllTasks() ([]*DTOs.TaskDTO, error) {
	tasks, err := repositories.GetAllTasks()
	if err != nil {
		return nil, err
	}

	var tasksDTOs []*DTOs.TaskDTO
	for _, task := range tasks {
		taskDTO := &DTOs.TaskDTO{}
		err = copier.Copy(taskDTO, task)
		if err != nil {
			log.Printf("Erreur lors de la copie des champs: %v", err)
			continue
		}
		tasksDTOs = append(tasksDTOs, taskDTO)
	}

	return tasksDTOs, nil
}

// Fonction qui retourne une tâche par son id
// Convertion de la tâche DAOs en tâche DTOs
func GetTaskById(id string) (*DTOs.TaskDTO, error) {
	task, err := repositories.GetTaskById(id)
	if err != nil {
		return nil, err
	} else if task == nil {
		return nil, nil
	}

	taskDTO := &DTOs.TaskDTO{}
	err = copier.Copy(taskDTO, task)
	if err != nil {
		log.Printf("Erreur lors de la copie des champs: %v", err)
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
