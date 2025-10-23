package models

import "time"

type TeamMember struct {
	UserID   uint64    `gorm:"column:user_id;primaryKey"`
	TeamID   uint64    `gorm:"column:team_id;primaryKey"`
	Role     string    `gorm:"size:50;not null"`
	JoinedAt time.Time `gorm:"column:joined_at"`

	// 외래키 설정
	User *User `gorm:"foreignKey:UserID;references:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Team *Team `gorm:"foreignKey:TeamID;references:TeamID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (TeamMember) TableName() string { return "team_members" }
