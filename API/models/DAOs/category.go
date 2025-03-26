package DAOs

import (
	"time"
)

type Category struct {
	Id          int        `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	Name        string     `json:"name" gorm:"type:varchar(50);not null"`
	Description string     `json:"description" gorm:"type:text;not null"`
	Activities  []Activity `json:"activities" gorm:"foreignKey:CategoryId;references:Id"`
	CreatedAt   time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time  `json:"updated_at" gorm:"autoUpdateTime"`
}
