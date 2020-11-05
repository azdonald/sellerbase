package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

//Base is the base struct for all models in this application
type Base struct {
	ID        string     `sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"update_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}

//BeforeSave gorm hook
func (base *Base) BeforeSave(tx *gorm.DB) (err error) {
	uuid := uuid.NewV4().String()

	base.ID = uuid
	return
}
