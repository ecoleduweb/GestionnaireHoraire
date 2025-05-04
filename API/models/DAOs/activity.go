package DAOs

import (
	"time"

	"gorm.io/gorm"
)

type Activity struct {
	Id          int       `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	Name        string    `json:"name" gorm:"type:varchar(50);not null"`
	Description string    `json:"description" gorm:"type:text;not null"`
	StartDate   time.Time `json:"start_date" gorm:"not null"`
	EndDate     time.Time `json:"end_date" gorm:"not null"`
	TimeSpent   float64   `json:"time_spent" gorm:"-"`
	// Clés étrangères
	UserId     int `json:"userId" gorm:"not null"`
	ProjectId  int `json:"projectId" gorm:"not null"`
	CategoryId int `json:"categoryId" gorm:"not null"`

	// Relations
	User     User     `json:"user" gorm:"foreignKey:UserId;references:Id"`
	Project  Project  `json:"project" gorm:"foreignKey:ProjectId;references:Id"`
	Category Category `json:"category" gorm:"foreignKey:CategoryId;references:Id"`
}

type ActivityWithTimeSpent struct {
	UserID     int     `gorm:"column:user_id"`
	CategoryID int     `gorm:"column:category_id"`
	ProjectID  int     `gorm:"column:project_id"`
	TimeSpent  float64 `gorm:"column:time_spent"`
}

func (a *Activity) AfterFind(tx *gorm.DB) error {
	// Only calculate if not already set (by aggregation query)
	if a.TimeSpent == 0 && !a.StartDate.IsZero() && !a.EndDate.IsZero() {
		a.TimeSpent = a.EndDate.Sub(a.StartDate).Hours()
	}
	return nil
}
