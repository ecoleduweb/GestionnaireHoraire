package repositories

import (
	"llio-api/database"
	"llio-api/models/DAOs"
)

func CreateActivity(activity *DAOs.Activity) (*DAOs.Activity, error) {
	err := database.DB.Create(activity).Error
	return activity, DBErrorManager(err)
}

func GetUsersActivities(userId int) ([]*DAOs.Activity, error) {
	var activities []*DAOs.Activity

	err := database.DB.Where("user_id = ?", userId).Find(&activities).Error
	return activities, DBErrorManager(err)
}

func GetActivityById(id string) (*DAOs.Activity, error) {
	var activity DAOs.Activity

	err := database.DB.First(&activity, id).Error
	return &activity, DBErrorManager(err)
}

func UpdateActivity(ActivityDAO *DAOs.Activity) (*DAOs.Activity, error) {
	err := database.DB.Updates(ActivityDAO).Error
	return ActivityDAO, DBErrorManager(err)
}

func DeleteActivity(id string) error {
	return DBErrorManager(database.DB.Delete(&DAOs.Activity{}, id).Error)
}

func GetActivitiesFromRange(from, to, idUser string) ([]*DAOs.Activity, error) {
	var activities []*DAOs.Activity
	err := database.DB.Where("Start_Date >= ? AND End_Date <= ? AND User_Id = ?", from, to, idUser).Find(&activities).Error
	return activities, DBErrorManager(err)
}
