package repositories

import (
	"llio-api/database"
	"llio-api/models/DAOs"
)

func CreateTask(task *DAOs.Task) (*DAOs.Task, error) {
	err := database.DB.Create(task).Error
	return task, err
}

func GetAllTasks() ([]*DAOs.Task, error) {
	var tasks []*DAOs.Task
	err := database.DB.Find(&tasks).Error
	return tasks, err
}

func GetTaskById(id string) (*DAOs.Task, error) {
	var task DAOs.Task

	err := database.DB.First(&task, id).Error
	return &task, err
}

func UpdateTask(taskDAO *DAOs.Task) (*DAOs.Task, error) {
	err := database.DB.Updates(taskDAO).Error
	return taskDAO, err
}

func DeleteTask(id string) error {
	return database.DB.Delete(&DAOs.Task{}, id).Error
}
