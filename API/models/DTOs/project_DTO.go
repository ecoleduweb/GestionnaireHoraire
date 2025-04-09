package DTOs

import (
	"llio-api/models/enums"
	"time"
)

type ProjectDTO struct {
	Id          int                 `json:"id"`
	Name        string              `json:"name" binding:"required,max=50"`
	Description string              `json:"description" binding:"required"`
	Status      enums.ProjectStatus `json:"status"` // Status du projet( si le projet est en cours ou fini)
	Billable    bool                `json:"billable"`
	Activities  []ActivityDTO       `json:"activities"`
	Categories  []CategoryDTO       `json:"categories"`
	CreatedAt   time.Time           `json:"created_at"`
	UpdatedAt   time.Time           `json:"updated_at"`
	EndAt       time.Time           `json:"end_at"`
}
