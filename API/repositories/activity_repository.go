package repositories

import (
	"llio-api/database"
	"llio-api/models/DAOs"
)

func CreateActivity(activities *DAOs.Activity) error {
	return database.DB.Create(activities).Error
}
