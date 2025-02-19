package repositories

import (
	"llio-api/database"
	"llio-api/models/DAOs"
)

// CreatTask insère une nouvelle tâche dans la BD
func CreateTask(task *DAOs.Task) error {
	return database.DB.Create(task).Error
}
