package repositories

import (
	"llio-api/database"
	"llio-api/models/DAOs"
)

// CreatTask insère une nouvelle tâche dans la BD
func CreateTask(task *DAOs.Task) error {
	return database.DB.Create(task).Error
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
	return &task, err
}
