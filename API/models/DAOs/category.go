package DAOs

import (
	"time"
)

type Category struct {
	Id          int        `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	Name        string     `json:"name" gorm:"type:varchar(50);not null; uniqueIndex:idx_name_project"`
	Description string     `json:"description" gorm:"type:text;not null"`
	Activities  []Activity `json:"activities" gorm:"foreignKey:CategoryId;references:Id"`
	CreatedAt   time.Time  `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt   time.Time  `json:"updatedAt" gorm:"autoUpdateTime"`

	// Clés étrangères
	UserId    int `json:"userId" gorm:"not null"`
	ProjectId int `json:"projectId" gorm:"uniqueIndex:idx_name_project"`

	// Relations
	User    User    `json:"user" gorm:"foreignKey:UserId;references:Id"`
	Project Project `json:"project" gorm:"foreignKey:ProjectId;references:Id"`
}
