package models

import (
	"time"

	"gorm.io/datatypes"
)

type Policy struct {
	PolicyID    uint64         `gorm:"column:policy_id;primaryKey;autoIncrement"`
	TeamID      uint64         `gorm:"column:team_id;not null;index"`
	PolicyName  string         `gorm:"column:policy_name;size:200;not null"`
	Description string         `gorm:"type:text"`
	Version     string         `gorm:"size:50"`
	SchemaA     datatypes.JSON `gorm:"column:schema_a_content;type:jsonb;not null"`
	SchemaB     datatypes.JSON `gorm:"column:schema_b_content;type:jsonb"`
	IsPublic    bool           `gorm:"column:is_public"`
	CreatedAt   time.Time      `gorm:"column:created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at"`

	Team *Team `gorm:"foreignKey:TeamID;references:TeamID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (Policy) TableName() string { return "policies" }
