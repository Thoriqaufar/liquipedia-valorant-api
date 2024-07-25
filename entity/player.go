package entity

import "time"

type Player struct {
	ID          uint      `gorm:"column:id;primaryKey"`
	TeamID      uint      `gorm:"column:team_id"`
	Name        string    `gorm:"column:name"`
	Nickname    string    `gorm:"column:nickname"`
	Nationality string    `gorm:"column:nationality"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	Team        Team      `gorm:"foreignKey:team_id;references:id"`
}
