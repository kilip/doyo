package orm

import (
	"github.com/kilip/doyo/orm/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// ORM interface
type ORM struct {
	DB *gorm.DB
}

// Open connection to sqlite
func (o ORM) Open(path string) error {
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})

	if err != nil {
		return err
	}

	db.AutoMigrate(&models.Project{})

	return nil
}
