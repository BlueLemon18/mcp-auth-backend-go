package models

import "time"

type ProjectPolicy struct {
	ProjectID uint64    `gorm:"column:project_id;primaryKey"`
	PolicyID  uint64    `gorm:"column:policy_id;primaryKey"`
	IsActive  bool      `gorm:"column:is_active;default:true"`
	LinkedAt  time.Time `gorm:"column:linked_at"`

	Project *Project `gorm:"foreignKey:ProjectID;references:ProjectID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Policy  *Policy  `gorm:"foreignKey:PolicyID;references:PolicyID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (ProjectPolicy) TableName() string { return "project_policies" }
