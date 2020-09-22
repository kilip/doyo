package models

import "gorm.io/gorm"

// Command struct
type Command struct {
	gorm.Model
	ProjectID      string
	Executable     string
	ExecutablePath string
}
