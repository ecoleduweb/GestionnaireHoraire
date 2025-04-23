package repositories

import (
	"llio-api/database"
	"llio-api/models/DAOs"
	"llio-api/structs"
)

func CreateActivity(activity *DAOs.Activity) (*structs.ActivityJoin, error) {
	err := database.DB.Create(activity).Error
	if err != nil {
		return nil, DBErrorManager(err)
	}

	var createdActivity structs.ActivityJoin
	err = database.DB.Table("activities").
		Select("activities.*, projects.name as project_name").
		Joins("LEFT JOIN projects ON activities.project_id = projects.id").
		Where("activities.id = ?", activity.Id).
		Scan(&createdActivity).Error

	return &createdActivity, DBErrorManager(err)
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

func UpdateActivity(ActivityDAO *DAOs.Activity) (*structs.ActivityJoin, error) {
	err := database.DB.Updates(ActivityDAO).Error
	if err != nil {
		return nil, DBErrorManager(err)
	}

	var updatedActivity structs.ActivityJoin
	err = database.DB.Table("activities").
		Select("activities.*, projects.name as project_name").
		Joins("LEFT JOIN projects ON activities.project_id = projects.id").
		Where("activities.id = ?", ActivityDAO.Id).
		Scan(&updatedActivity).Error

	return &updatedActivity, DBErrorManager(err)
}

func DeleteActivity(id string) error {
	return DBErrorManager(database.DB.Delete(&DAOs.Activity{}, id).Error)
}

func GetActivitiesFromRange(from, to, idUser string) ([]*structs.ActivityJoin, error) {
	var activities []*structs.ActivityJoin

	err := database.DB.Table("activities").
		Select("activities.*, projects.name as project_name").
		Joins("LEFT JOIN projects ON activities.project_id = projects.id").
		Where("Start_Date >= ? AND End_Date <= ? AND User_Id = ?", from, to, idUser).
		Scan(&activities).Error

	return activities, DBErrorManager(err)
}
