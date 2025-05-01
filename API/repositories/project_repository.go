package repositories

import (
	"llio-api/database"
	"llio-api/models/DAOs"
)

func CreateProject(project *DAOs.Project) (*DAOs.Project, error) {
	err := database.DB.Create(project).Error
	return project, DBErrorManager(err)
}

func GetProjects() ([]*DAOs.Project, error) {
	var projects []*DAOs.Project
	err := database.DB.Preload("Activities").Preload("Activities.User").Preload("Activities.Category").Preload("Categories").Find(&projects).Error
	return projects, DBErrorManager(err)
}

func GetProjectsByManagerId(id int) ([]*DAOs.Project, error) {
	var projects []*DAOs.Project
	err := database.DB.Preload("Activities").Preload("Activities.User").Preload("Activities.Category").Preload("Categories").Find(&projects, "manager_id = ?", id).Error
	return projects, DBErrorManager(err)
}

func GetProjectById(id string) (*DAOs.Project, error) {
	var project DAOs.Project

	err := database.DB.First(&project, id).Error
	return &project, DBErrorManager(err)
}

func UpdateProject(projectDAO *DAOs.Project) (*DAOs.Project, error) {
	err := database.DB.Updates(projectDAO).Error
	return projectDAO, DBErrorManager(err)
}
