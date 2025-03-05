package DTOs

import "time"

// TaskRequest représente les données entrantes pour créer une tâche
type TaskDTO struct {
	Id          int       `json:"id"`
	Name        string    `json:"name" binding:"required,max=50"`
	Description string    `json:"description" binding:"required"`
	Billable    bool      `json:"billable"`
	StartDate   time.Time `json:"startDate" binding:"required"`
	EndDate     time.Time `json:"endDate" binding:"required"`
	UserId      int       `json:"userId" `
	ProjectId   int       `json:"projectId" `
	CategoryId  int       `json:"categoryId" `
}
