package user

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"size:100;not null" json:"name"`
	Email     string    `gorm:"size:255;unique;not null" json:"email"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	IsActive  bool      `gorm:"default:true" json:"is_active"`
}
