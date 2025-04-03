package DTOs

import (
	"time"
)

type CategoryDTO struct {
	Id          int           `json:"id" `
	Name        string        `json:"name" binding:"required,max=50"`
	Description string        `json:"description" binding:"required"`
	Billable    bool          `json:"billable"`
	Activities  []ActivityDTO `json:"activities" `
	CreatedAt   time.Time     `json:"createdAt" `
	UpdatedAt   time.Time     `json:"updatedAt" `
	UserId      int           `json:"userId" `
	ProjectId   int           `json:"projectId" `
}
