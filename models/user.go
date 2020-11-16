package models

import (
	"errors"

	"github.com/goware/emailx"
)

// User represents an individual in the system.
// Users belong to a merchant
type User struct {
	Base
	FirstName  string   `gorm:"size:100;not null" json:"first_name"`
	LastName   string   `gorm:"size:100;not null" json:"last_name"`
	Email      string   `gorm:"size:100;not null;unique" json:"email"`
	MerchantID string   `sql:"type:uuid" gorm:"size:100;" json:"merchant_id"`
	Merchant   Merchant `json:"-"`
}

//IsValid checks if required params are valid
func (u *User) IsValid() (bool, error) {
	err := emailx.Validate(u.Email)
	if err != nil {
		return false, errors.New("Email is not valid")
	}

	if u.FirstName == "" {
		return false, errors.New("First name cannot be empty")
	}

	if u.LastName == "" {
		return false, errors.New("Last name cannot be empty")
	}

	return true, nil
}
