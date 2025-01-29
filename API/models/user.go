package models

type User struct {
	Id        int    `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	FirstName string `json:"first_name" gorm:"type:varchar(50);not null"`
	LastName  string `json:"last_name" gorm:"type:varchar(50);not null"`
	Email     string `json:"email" gorm:"type:varchar(255);not null;unique;index"`
	Role      string `json:"role" gorm:"type:enum('gestionnaire de projet','chargé de projet','employé');not null;default:'employé'"`
	Tasks     []Task `json:"tasks" gorm:"foreignKey:UserId;references:Id"`
}
