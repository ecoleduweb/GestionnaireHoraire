package services

import (
	"llio-api/models/DAOs"
	"llio-api/models/DTOs"
	"llio-api/repositories"

	"github.com/jinzhu/copier"
)

func VerifyCreateCategoryJSON(categoryDTO *DTOs.CategoryDTO) []DTOs.FieldErrorDTO {
	var errors []DTOs.FieldErrorDTO

	// Vérifier que StartDate est avant EndDate
	if categoryDTO.CreatedAt.After(categoryDTO.UpdatedAt) {
		errors = append(errors, DTOs.FieldErrorDTO{
			Field:   "createdAt",
			Message: "La date de création doit être avant la date de mise à jour",
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
