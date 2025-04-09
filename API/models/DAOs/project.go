package DAOs

import (
	"llio-api/models/enums"
	"time"
)

type Project struct {
	Id          int                 `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	Name        string              `json:"name" gorm:"type:varchar(50);not null;unique"`
	Description string              `json:"description" gorm:"type:text;not null"`
	Status      enums.ProjectStatus `json:"status" gorm:"type:enum(3,2,1,0);not null;default:0"`
	Billable    bool                `json:"billable" gorm:"type:boolean;not null;default:false"`
	Activities  []Activity          `json:"activities" gorm:"foreignKey:ProjectId;references:Id"`
	Categories  []Category          `json:"categories" gorm:"foreignKey:ProjectId;references:Id"`
	CreatedAt   time.Time           `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time           `json:"updated_at" gorm:"autoUpdateTime"`
	EndAt       time.Time           `json:"end_at"`
}
