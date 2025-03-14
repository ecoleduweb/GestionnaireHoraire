package services

import (
	"llio-api/models/DAOs"
	"llio-api/models/DTOs"
	"llio-api/repositories"
)

func VerifyCreateActivityJSON(activityDTO *DTOs.ActivityDTO) []DTOs.FieldErrorDTO {
	var errors []DTOs.FieldErrorDTO

	// Vérifier que StartDate est avant EndDate
	if activityDTO.StartDate.After(activityDTO.EndDate) {
		errors = append(errors, DTOs.FieldErrorDTO{
			Field:   "start_date",
			Message: "La date de début doit être avant la date de fin",
		})
	}

	return errors
}

// CreateActivityService encapsule la logique métier pour la création d'une tâche
func CreateActivityService(activityDTO *DTOs.ActivityDTO) (*DTOs.ActivityResponse, error) {

	activity := &DAOs.Activity{
		Name:        activityDTO.Name,
		Description: activityDTO.Description,
		StartDate:   activityDTO.StartDate,
		EndDate:     activityDTO.EndDate,
		UserId:      activityDTO.UserId,
		ProjectId:   activityDTO.ProjectId,
		CategoryId:  activityDTO.CategoryId,
	}

	err := repositories.CreateActivity(activity)
	if err != nil {
		return nil, err
	}

	activityResponse := &DTOs.ActivityResponse{
		Name:        activity.Name,
		Description: activity.Description,
		StartDate:   activity.StartDate,
		EndDate:     activity.EndDate,
	}
	return activityResponse, nil
}
