package util

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

//BeforeCreate gorm hook
func (base *Base) BeforeCreate(tx *gorm.DB) (err error) {
	uuid := uuid.NewV4().String()

	base.ID = uuid
	return
}
