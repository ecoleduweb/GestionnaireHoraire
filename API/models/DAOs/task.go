package DAOs

import (
	"time"
)

type Task struct {
	Id          int       `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	Name        string    `json:"name" gorm:"type:varchar(50);not null"`
	Description string    `json:"description" gorm:"type:text;not null"`
	State       string    `json:"state" gorm:"type:enum('à faire','en cours','terminé');not null;default:'à faire'"`
	Billable    bool      `json:"billable" gorm:"type:boolean;not null;default:false"`
	StartDate   time.Time `json:"start_date" gorm:"not null"`
	EndDate     time.Time `json:"end_date" gorm:"not null"`

	// Clés étrangères
	UserId     int `json:"user_id" gorm:"not null"`
	ProjectId  int `json:"project_id" gorm:"not null"`
	CategoryId int `json:"category_id" gorm:"not null"`

	// Relations
	User     User     `json:"user" gorm:"foreignKey:UserId;references:Id"`
	Project  Project  `json:"project" gorm:"foreignKey:ProjectId;references:Id"`
	Category Category `json:"category" gorm:"foreignKey:CategoryId;references:Id"`
}
