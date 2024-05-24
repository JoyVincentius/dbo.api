// models/auth.go
package models

import "time"

type User struct {
	ID        uint   `gorm:"primaryKey" json:"-"`
	Username  string `gorm:"size:255;not null;unique" json:"username"`
	Password  string `gorm:"size:255;not null" json:"password"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
