package models

import (
	"time"
)

// -(2025-01-29/Quentin) : Version de ChatGPT en attendant version de Lyes
type Task struct {
	ID          uint      `gorm:"primaryKey"`
	UserID      uint      `gorm:"not null"`
	IdProject   uint      `gorm:"not null"`
	IdCategory  uint      `gorm:"not null"`
	Name        string    `gorm:"size:50;not null"`
	Description string    `gorm:"type:text"`
	State       string    `gorm:"type:enum('À faire', 'En cours', 'Terminé');default:'À faire'"`
	Billable    bool      `gorm:"default:false"`
	StartDate   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	EndAt       time.Time
}
