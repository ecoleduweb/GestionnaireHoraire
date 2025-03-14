package repositories

import (
	"llio-api/database"
	"llio-api/models/DAOs"
)

func CreateActivity(activity *DAOs.Activity) (*DAOs.Activity, error) {
	err := database.DB.Create(activity).Error
	return activity, err
}

func GetAllActivities() ([]*DAOs.Activity, error) {
	var activities []*DAOs.Activity
	err := database.DB.Find(&activities).Error
	return activities, err
}

func GetActivityById(id string) (*DAOs.Activity, error) {
	var activity DAOs.Activity

	err := database.DB.First(&activity, id).Error
	return &activity, err
}

func UpdateActivity(ActivityDAO *DAOs.Activity) (*DAOs.Activity, error) {
	err := database.DB.Updates(ActivityDAO).Error
	return ActivityDAO, err
}
