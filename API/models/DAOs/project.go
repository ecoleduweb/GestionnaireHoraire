package DAOs

import (
	"time"
)

type Project struct {
	Id          int       `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	Name        string    `json:"name" gorm:"type:varchar(50);not null"`
	Description string    `json:"description" gorm:"type:text;not null"`
	Status      bool      `json:"status" gorm:"type:boolean;not null;default:true"`
	Task        []Task    `json:"task" gorm:"foreignKey:ProjectId;references:Id"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	EndAt       time.Time `json:"end_at"`
}
