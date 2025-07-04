package DAOs

import (
	"llio-api/models/enums"
	"time"
)

type Project struct {
	Id             int                 `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	ManagerId      int                 `json:"manager_id" gorm:"type:bigint(20) unsigned;not null"`
	UniqueId       string              `json:"uniqueId" gorm:"type:varchar(20);not null;unique"`
	Name           string              `json:"name" gorm:"type:text;not null"`
	Status         enums.ProjectStatus `json:"status" gorm:"type:enum(3,2,1,0);not null;default:0"` // Status du projet( si le projet est en cours ou fini)
	Billable       bool                `json:"billable" gorm:"type:boolean;not null;default:false"`
	Activities     []Activity          `json:"activities" gorm:"foreignKey:ProjectId;references:Id"`
	Categories     []Category          `json:"categories" gorm:"foreignKey:ProjectId;references:Id"`
	CreatedAt      time.Time           `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt      time.Time           `json:"updated_at" gorm:"autoUpdateTime"`
	EndAt          time.Time           `json:"end_at"`
	EstimatedHours int                 `json:"estimated_hours" gorm:"type:int(11);not null;default:0"` // Heures estimées pour le projet
}
