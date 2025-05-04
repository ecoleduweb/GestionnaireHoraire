package DAOs

import (
	"time"
)

type Activity struct {
	Id          int       `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	Name        string    `json:"name" gorm:"type:varchar(50);not null"`
	Description string    `json:"description" gorm:"type:text;not null"`
	StartDate   time.Time `json:"start_date" gorm:"not null"`
	EndDate     time.Time `json:"end_date" gorm:"not null"`
	TimeSpent   float64   `json:"time_spent" gorm:"type:float"`
	// Clés étrangères
	UserId     int `json:"userId" gorm:"not null"`
	ProjectId  int `json:"projectId" gorm:"not null"`
	CategoryId int `json:"categoryId" gorm:"not null"`

	// Relations
	User     User     `json:"user" gorm:"foreignKey:UserId;references:Id"`
	Project  Project  `json:"project" gorm:"foreignKey:ProjectId;references:Id"`
	Category Category `json:"category" gorm:"foreignKey:CategoryId;references:Id"`
}
