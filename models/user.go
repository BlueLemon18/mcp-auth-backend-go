package models

import "time"

type User struct {
	UserID       uint64    `gorm:"column:user_id;primaryKey;autoIncrement"`
	Username     string    `gorm:"size:100;not null"`
	Email        string    `gorm:"size:255;not null;uniqueIndex"`
	PasswordHash string    `gorm:"size:255;not null"`
	CreatedAt    time.Time `gorm:"column:created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at"`
}

func (User) TableName() string { return "users" }
