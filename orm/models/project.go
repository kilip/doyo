package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// Project struct
// gorm.Model definition
type Project struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
	Path      string    `gorm:"type:string;"`
	Name      string    `gorm:"type:string;"`
	Commands  []Command `gorm:"foreignKey:ProjectID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
