package repositories

import (
	"errors"
	"llio-api/database"
	"llio-api/models/DAOs"

	"gorm.io/gorm"
)

// CreatTask insère une nouvelle tâche dans la BD
func CreateTask(task *DAOs.Task) (*DAOs.Task, error) {
	err := database.DB.Create(task).Error
	if err != nil {
		return nil, err
	}
	return task, nil
}

// GetAllTasks retourne toutes les tâches de la BD
func GetAllTasks() ([]*DAOs.Task, error) {
	var tasks []*DAOs.Task
	err := database.DB.Find(&tasks).Error
	return tasks, err
}

// GetTaskById retourne une tâche par son id
func GetTaskById(id string) (*DAOs.Task, error) {
	var task DAOs.Task

	err := database.DB.First(&task, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		// Une autre erreur s'est produite
		return nil, err
	}
	return &task, err
}

// Fonction qui permet de mettre à jour une tâche
func UpdateTask(taskDAO *DAOs.Task) (*DAOs.Task, error) {
	err := database.DB.Updates(taskDAO).Error
	if err != nil {
		return nil, err
	}
	return taskDAO, nil
}

// Fonction qui permet la suppression de la tâche dans la BD
func DeleteTask(id string) error {
	return database.DB.Delete(&DAOs.Task{}, id).Error
}
