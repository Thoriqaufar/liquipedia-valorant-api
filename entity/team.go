package entity

import (
	"time"
)

type Team struct {
	ID        uint      `gorm:"column:id;primaryKey"`
	Name      string    `gorm:"column:name"`
	Alias     string    `gorm:"column:alias"`
	Origin    string    `gorm:"column:origin"`
	Region    string    `gorm:"column:region"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	Player    []Player  `gorm:"foreignKey:team_id;references:id"`
}
