package repositories

import (
	"errors"
	"llio-api/database"
	"llio-api/models/DAOs"

	"gorm.io/gorm"
)

func CreateTask(task *DAOs.Task) (*DAOs.Task, error) {
	err := database.DB.Create(task).Error
	if err != nil {
		return nil, err
	}
	return task, nil
}

func GetAllTasks() ([]*DAOs.Task, error) {
	var tasks []*DAOs.Task
	err := database.DB.Find(&tasks).Error
	return tasks, err
}

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
