package DTOs

import (
	"llio-api/models/enums"
	"time"
)

type ProjectDTO struct {
	Id             int                 `json:"id"`
	ManagerId      int                 `json:"managerId" binding:"required"`
	UniqueId       string              `json:"uniqueId" binding:"required,max=50"`
	Name           string              `json:"name"`
	Status         enums.ProjectStatus `json:"status"` // Status du projet( si le projet est en cours ou fini)
	Billable       bool                `json:"billable"`
	Activities     []ActivityDTO       `json:"activities"`
	Categories     []CategoryDTO       `json:"categories"`
	CreatedAt      time.Time           `json:"createdAt"`
	UpdatedAt      time.Time           `json:"updatedAt"`
	EndAt          time.Time           `json:"endAt"`
	EstimatedHours int                 `json:"estimatedHours"`
}
