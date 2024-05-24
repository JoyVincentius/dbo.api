// models/customer.go
package models

import "time"

type Customer struct {
	ID        uint   `gorm:"primaryKey" json:"-"`
	Name      string `gorm:"size:255;not null" json:"name"`
	Email     string `gorm:"size:255;not null;unique" json:"email"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
