package services

import (
	"llio-api/models/DAOs"
	"llio-api/models/DTOs"
	"llio-api/repositories"

	"github.com/jinzhu/copier"
)

func VerifyCreateTaskJSON(taskDTO *DTOs.TaskDTO) []DTOs.FieldErrorDTO {
	var errors []DTOs.FieldErrorDTO

	if taskDTO.UserId == 0 {
		errors = append(errors, DTOs.FieldErrorDTO{
			Field:   "user_id",
			Message: "Le champ UserId est invalide ou manquant",
		})
	}
	if taskDTO.ProjectId == 0 {
		errors = append(errors, DTOs.FieldErrorDTO{
			Field:   "project_id",
			Message: "Le champ ProjectId est invalide ou manquant",
		})
	}
	if taskDTO.CategoryId == 0 {
		errors = append(errors, DTOs.FieldErrorDTO{
			Field:   "category_id",
			Message: "Le champ CategoryId est invalide ou manquant",
		})
	}
	// Vérifier que StartDate est avant EndDate
	if taskDTO.StartDate.After(taskDTO.EndDate) {
		errors = append(errors, DTOs.FieldErrorDTO{
			Field:   "start_date",
			Message: "La date de début doit être avant la date de fin",
		})
	}

	return errors
}

func CreateTask(taskDTO *DTOs.TaskDTO) (*DTOs.TaskDTO, error) {

	// Mapper le body vers le modèle Task

	task := &DAOs.Task{}
	err := copier.Copy(task, taskDTO)
	if err != nil {
		return nil, err
	}

	taskDAOAded, err := repositories.CreateTask(task)
	if err != nil {
		return nil, err
	}

	taskDTOResponse := &DTOs.TaskDTO{}
	err = copier.Copy(taskDTOResponse, taskDAOAded)
	return taskDTOResponse, err
}

func GetAllTasks() ([]*DTOs.TaskDTO, error) {
	tasks, err := repositories.GetAllTasks()
	if err != nil {
		return nil, err
	}

	var tasksDTOs []*DTOs.TaskDTO
	for _, task := range tasks {
		taskDTO := &DTOs.TaskDTO{}
		err = copier.Copy(taskDTO, task)
		tasksDTOs = append(tasksDTOs, taskDTO)
	}

	return tasksDTOs, err
}

func GetTaskById(id string) (*DTOs.TaskDTO, error) {
	task, err := repositories.GetTaskById(id)
	if err != nil {
		return nil, err
	}

	taskDTO := &DTOs.TaskDTO{}
	err = copier.Copy(taskDTO, task)

	return taskDTO, err
}

func UpdateTask(taskDTO *DTOs.TaskDTO) (*DTOs.TaskDTO, error) {

	taskDAO := &DAOs.Task{}
	err := copier.Copy(taskDAO, taskDTO)
	if err != nil {
		log.Printf("Erreur lors de la copie des champs: %v", err)
	}

	taskDAOUpdated, err := repositories.UpdateTask(taskDAO)
	if err != nil {
		return nil, err
	}

	taskDTOResponse := &DTOs.TaskDTO{}
	err = copier.Copy(taskDTOResponse, taskDAOUpdated)
	if err != nil {
		log.Printf("Erreur lors de la copie des champs: %v", err)
		return nil, err
	}
	return taskDTOResponse, nil
}
