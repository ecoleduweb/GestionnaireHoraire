package repositories

import (
	"llio-api/database"
	"llio-api/models/DAOs"
)

func CreateCategory(category *DAOs.Category) (*DAOs.Category, error) {
	err := database.DB.Create(category).Error
	return category, err
}

func GetCategories() ([]*DAOs.Category, error) {
	var categories []*DAOs.Category
	err := database.DB.Find(&categories).Error
	return categories, err
}

func GetCategoryById(id string) (*DAOs.Category, error) {
	var category DAOs.Category

	err := database.DB.First(&category, id).Error
	return &category, err
}

func UpdateCategory(categoryDAO *DAOs.Category) (*DAOs.Category, error) {
	err := database.DB.Updates(categoryDAO).Error
	return categoryDAO, err
}

func DeleteCategory(id string) error {
	return database.DB.Delete(&DAOs.Category{}, id).Error
}
