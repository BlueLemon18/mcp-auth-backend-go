package models

import "time"

type Team struct {
	TeamID      uint64    `gorm:"column:team_id;primaryKey;autoIncrement"`
	TeamName    string    `gorm:"size:150;not null"`
	OwnerUserID uint64    `gorm:"column:owner_user_id;not null;index"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`

	// 관계
	Owner *User `gorm:"foreignKey:OwnerUserID;references:UserID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
}

func (Team) TableName() string { return "teams" }
