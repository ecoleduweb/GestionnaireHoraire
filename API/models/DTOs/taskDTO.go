package DTOs

import "time"

// TaskRequest représente les données entrantes pour créer une tâche
type TaskDTO struct {
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	State       string    `json:"state"`
	Billable    bool      `json:"billable"`
	StartDate   time.Time `json:"start_date" binding:"required"`
	EndDate     time.Time `json:"end_date" binding:"required"`
	UserId      int       `json:"user_id" binding:"required"`
	ProjectId   int       `json:"project_id" binding:"required"`
	CategoryId  int       `json:"category_id" binding:"required"`
}
