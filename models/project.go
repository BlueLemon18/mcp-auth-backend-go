package models

import "time"

type Project struct {
	ProjectID     uint64    `gorm:"column:project_id;primaryKey;autoIncrement"`
	TeamID        uint64    `gorm:"column:team_id;not null;index"`
	ProjectName   string    `gorm:"column:project_name;size:200;not null"`
	RepositoryURL string    `gorm:"column:repository_url;size:500"`
	CreatedAt     time.Time `gorm:"column:created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at"`

	Team *Team `gorm:"foreignKey:TeamID;references:TeamID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (Project) TableName() string { return "projects" }
