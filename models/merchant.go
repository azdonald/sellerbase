package models

// Merchant represents the sellers on the platform
type Merchant struct {
	Base
	BusinessName string `gorm:"size:100;not null;unique" json:"business_name"`
}
