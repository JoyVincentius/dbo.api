// models/order.go
package models

import "time"

type Order struct {
	ID         uint   `gorm:"primaryKey" json:"-"`
	CustomerID uint   `gorm:"not null" json:"customer_id"`
	Item       string `gorm:"size:255;not null" json:"item"`
	Quantity   int    `gorm:"not null" json:"quantity"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
