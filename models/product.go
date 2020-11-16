package models

// Product represents an item
type Product struct {
	Base
	Name        string  `gorm:"size:100;not null" json:"name"`
	Description string  `gorm:"size:250;not null" json:"description"`
	SKU         string  `gorm:"size:100;not null" json:"sku"`
	Price       float64 `gorm:"not null" json:"price"`
	MerchantID  string  `sql:"type:uuid" gorm:"size:100;" json:"merchant_id"`
}
