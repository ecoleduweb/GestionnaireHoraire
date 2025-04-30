package services

import (
	"llio-api/models/DAOs"
	"llio-api/models/DTOs"
	"llio-api/repositories"
	"log"

	"github.com/jinzhu/copier"
)

func VerifyCreateActivityJSON(activityDTO *DTOs.ActivityDTO) []DTOs.FieldErrorDTO {
	var errors []DTOs.FieldErrorDTO

	if activityDTO.UserId == 0 {
		errors = append(errors, DTOs.FieldErrorDTO{
			Field:   "userId",
			Message: "Le champ userId est invalide ou manquant",
		})
	}
	if activityDTO.ProjectId == 0 {
		errors = append(errors, DTOs.FieldErrorDTO{
			Field:   "projectId",
			Message: "Le champ projectId est invalide ou manquant",
		})
	}
	if activityDTO.CategoryId == 0 {
		errors = append(errors, DTOs.FieldErrorDTO{
			Field:   "categoryId",
			Message: "Le champ categoryId est invalide ou manquant",
		})
	}
	// Vérifier que StartDate est avant EndDate
	if activityDTO.StartDate.After(activityDTO.EndDate) {
		errors = append(errors, DTOs.FieldErrorDTO{
			Field:   "startDate",
			Message: "La date de début doit être avant la date de fin",
		})
	}

	return errors
}

func CreateActivity(activityDTO *DTOs.ActivityDTO) (*DTOs.ActivityInfoDTO, error) {

	// Mapper le body vers le modèle Activity

	activity := &DAOs.Activity{}
	err := copier.Copy(activity, activityDTO)
	if err != nil {
		return nil, err
	}

	activityDAOAded, err := repositories.CreateActivity(activity)
	if err != nil {
		return nil, err
	}

	activityDTOResponse := &DTOs.ActivityInfoDTO{}
	err = copier.Copy(activityDTOResponse, activityDAOAded)
	return activityDTOResponse, err
}

func GetUsersActivities(userId int) ([]*DTOs.ActivityDTO, error) {
	activities, err := repositories.GetUsersActivities(userId)
	if err != nil {
		return nil, err
	}

	var activitiesDTOs []*DTOs.ActivityDTO
	for _, activity := range activities {
		activityDTO := &DTOs.ActivityDTO{}
		err = copier.Copy(activityDTO, activity)
		activitiesDTOs = append(activitiesDTOs, activityDTO)
	}

	return activitiesDTOs, err
}

func GetActivityById(id string) (*DTOs.ActivityDTO, error) {
	activity, err := repositories.GetActivityById(id)
	if err != nil {
		return nil, err
	}

	activityDTO := &DTOs.ActivityDTO{}
	err = copier.Copy(activityDTO, activity)

	return activityDTO, err
}

func UpdateActivity(activityDTO *DTOs.ActivityDTO) (*DTOs.ActivityInfoDTO, error) {

	activityDAO := &DAOs.Activity{}
	err := copier.Copy(activityDAO, activityDTO)
	if err != nil {
		return nil, err
	}

	activityDAOUpdated, err := repositories.UpdateActivity(activityDAO)
	if err != nil {
		return nil, err
	}

	activityDTOResponse := &DTOs.ActivityInfoDTO{}
	err = copier.Copy(activityDTOResponse, activityDAOUpdated)
	return activityDTOResponse, err
}

func DeleteActivity(id string) error {
	return repositories.DeleteActivity(id)
}

func GetActivitiesFromRange(from, to, idUser string) ([]*DTOs.ActivityInfoDTO, error) {
	activities, err := repositories.GetActivitiesFromRange(from, to, idUser)
	if err != nil {
		return nil, err
	}

	var activitiesDTOs []*DTOs.ActivityInfoDTO
	for _, activity := range activities {
		activityDTO := &DTOs.ActivityInfoDTO{}
		err = copier.Copy(activityDTO, activity)
		if err != nil {
			return nil, err
		}
		activitiesDTOs = append(activitiesDTOs, activityDTO)
	}
	return activitiesDTOs, nil
}
