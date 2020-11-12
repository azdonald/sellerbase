package models

// User represents an individual in the system.
// Users belong to a merchant
type User struct {
	Base
	FirstName  string `gorm:"size:100;not null" json:"first_name"`
	LastName   string `gorm:"size:100;not null" json:"last_name"`
	Email      string `gorm:"size:100;not null;unique" json:"email"`
	MerchantID string `sql:"type:uuid" gorm:"size:100;" json:"merchant_id"`
	Merchant   Merchant
}
