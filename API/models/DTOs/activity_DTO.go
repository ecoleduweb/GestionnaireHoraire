package DTOs

import "time"

// TaskRequest représente les données entrantes pour créer une activités
type ActivityDTO struct {
	Id          int       `json:"id"`
	Name        string    `json:"name" binding:"max=50"`
	Description string    `json:"description"`
	StartDate   time.Time `json:"startDate" binding:"required"`
	EndDate     time.Time `json:"endDate" binding:"required"`
	UserId      int       `json:"userId" `
	ProjectId   int       `json:"projectId" `
	CategoryId  int       `json:"categoryId" `
}
