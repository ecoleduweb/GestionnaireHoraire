package DTOs

import "time"

// ActivityRequest représente les données entrantes pour créer une tâche
type ActivityDTO struct {
	Name        string    `json:"name" binding:"required,max=50"`
	Description string    `json:"description" binding:"required"`
	StartDate   time.Time `json:"start_date" binding:"required"`
	EndDate     time.Time `json:"end_date" binding:"required"`
	UserId      int       `json:"user_id" binding:"required"`
	ProjectId   int       `json:"project_id" binding:"required"`
	CategoryId  int       `json:"category_id" binding:"required"`
}

type ActivityResponse struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	State       string    `json:"state"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
}
