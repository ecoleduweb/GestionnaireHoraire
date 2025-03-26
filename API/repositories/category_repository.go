package repositories

import (
	"llio-api/database"
	"llio-api/models/DAOs"
)

func CreateCategory(category *DAOs.Category) (*DAOs.Category, error) {
	err := database.DB.Create(category).Error
	return category, err
}
