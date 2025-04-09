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
	var projets []*DAOs.Project
	err := database.DB.Find(&projets).Error
	return projets, DBErrorManager(err)
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
