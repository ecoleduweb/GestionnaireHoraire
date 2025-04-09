package services

import (
	"llio-api/models/DAOs"
	"llio-api/models/DTOs"
	"llio-api/repositories"

	"github.com/jinzhu/copier"
)

func VerifyProjectJSON(projectDTO *DTOs.ProjectDTO) []DTOs.FieldErrorDTO {
	var errors []DTOs.FieldErrorDTO

	if !projectDTO.CreatedAt.IsZero() {
		// Vérifier que CreatedAt est avant EndAt
		if !projectDTO.EndAt.IsZero() && projectDTO.CreatedAt.After(projectDTO.EndAt) {
			errors = append(errors, DTOs.FieldErrorDTO{
				Field:   "endAt",
				Message: "La date de fin doit être après la date de création",
			})
		}

		// Vérifier que CreatedAt est avant UpdatedAt
		if !projectDTO.UpdatedAt.IsZero() && projectDTO.CreatedAt.After(projectDTO.UpdatedAt) {
			errors = append(errors, DTOs.FieldErrorDTO{
				Field:   "updatedAt",
				Message: "La date de mise à jour doit être après la date de création",
			})
		}
	}

	return errors
}

func CreateProject(projectDTO *DTOs.ProjectDTO) (*DTOs.ProjectDTO, error) {

	project := &DAOs.Project{}
	err := copier.Copy(project, projectDTO)
	if err != nil {
		return nil, err
	}

	projectDAOAdded, err := repositories.CreateProject(project)
	if err != nil {
		return nil, err
	}

	projectDTOResponse := &DTOs.ProjectDTO{}
	err = copier.Copy(projectDTOResponse, projectDAOAdded)
	return projectDTOResponse, err
}

func GetProjects() ([]*DTOs.ProjectDTO, error) {
	projects, err := repositories.GetProjects()
	if err != nil {
		return nil, err
	}

	var projectsDTOs []*DTOs.ProjectDTO
	for _, project := range projects {
		projectDTO := &DTOs.ProjectDTO{}
		err = copier.Copy(projectDTO, project)
		projectsDTOs = append(projectsDTOs, projectDTO)
	}

	return projectsDTOs, err
}

func GetProjectById(id string) (*DTOs.ProjectDTO, error) {
	project, err := repositories.GetProjectById(id)
	if err != nil {
		return nil, err
	}

	projectDTO := &DTOs.ProjectDTO{}
	err = copier.Copy(projectDTO, project)

	return projectDTO, err
}

func UpdateProject(projectDTO *DTOs.ProjectDTO) (*DTOs.ProjectDTO, error) {

	projectDAO := &DAOs.Project{}
	err := copier.Copy(projectDAO, projectDTO)
	if err != nil {
		return nil, err
	}

	projectDAOUpdated, err := repositories.UpdateProject(projectDAO)
	if err != nil {
		return nil, err
	}

	projectDTOResponse := &DTOs.ProjectDTO{}
	err = copier.Copy(projectDTOResponse, projectDAOUpdated)
	return projectDTOResponse, err
}
