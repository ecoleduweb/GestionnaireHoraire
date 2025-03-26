package services

import (
	"llio-api/models/DAOs"
	"llio-api/models/DTOs"
	"llio-api/repositories"
	"log"

	"github.com/jinzhu/copier"
)

func VerifyCreateCategoryJSON(categoryDTO *DTOs.CategoryDTO) []DTOs.FieldErrorDTO {
	var errors []DTOs.FieldErrorDTO

	if categoryDTO.ProjectId == 0 {
		errors = append(errors, DTOs.FieldErrorDTO{
			Field:   "projectId",
			Message: "Le champ projectId est invalide ou manquant",
		})
	}

	return errors
}

func CreateCategory(categoryDTO *DTOs.CategoryDTO) (*DTOs.CategoryDTO, error) {

	// Mapper le body vers le modèle Activity

	category := &DAOs.Category{}
	err := copier.Copy(category, categoryDTO)
	if err != nil {
		return nil, err
	}

	activityDAOAded, err := repositories.CreateCategory(category)
	if err != nil {
		return nil, err
	}

	categoryDTOResponse := &DTOs.CategoryDTO{}
	err = copier.Copy(categoryDTOResponse, activityDAOAded)
	return categoryDTOResponse, err
}

func GetCategories() ([]*DTOs.CategoryDTO, error) {
	categories, err := repositories.GetCategories()
	if err != nil {
		return nil, err
	}

	var categoriesDTOs []*DTOs.CategoryDTO
	for _, category := range categories {
		categoryDTO := &DTOs.CategoryDTO{}
		err = copier.Copy(categoryDTO, category)
		categoriesDTOs = append(categoriesDTOs, categoryDTO)
	}

	return categoriesDTOs, err
}

func GetCategoryById(id string) (*DTOs.CategoryDTO, error) {
	category, err := repositories.GetCategoryById(id)
	if err != nil {
		return nil, err
	}

	categoryDTO := &DTOs.CategoryDTO{}
	err = copier.Copy(categoryDTO, category)

	return categoryDTO, err
}

func UpdateCategory(categoryDTO *DTOs.CategoryDTO) (*DTOs.CategoryDTO, error) {

	categoryDAO := &DAOs.Category{}
	err := copier.Copy(categoryDAO, categoryDTO)
	if err != nil {
		return nil, err
	}

	categoryDAOUpdated, err := repositories.UpdateCategory(categoryDAO)
	log.Println(categoryDAOUpdated)
	if err != nil {
		return nil, err
	}

	categoryDTOResponse := &DTOs.CategoryDTO{}
	err = copier.Copy(categoryDTOResponse, categoryDAOUpdated)
	return categoryDTOResponse, err
}
