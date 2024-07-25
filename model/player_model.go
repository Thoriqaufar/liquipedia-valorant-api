package model

import "time"

type PlayerResponse struct {
	ID          uint      `json:"id"`
	TeamID      uint      `json:"team_id"`
	Name        string    `json:"name"`
	Nickname    string    `json:"nickname"`
	Nationality string    `json:"nationality"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CreatePlayerRequest struct {
	TeamID      uint      `json:"team_id,string" validate:"required"`
	Name        string    `json:"name" validate:"required,max=255"`
	Nickname    string    `json:"nickname" validate:"required,max=255"`
	Nationality string    `json:"nationality" validate:"required,max=255"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UpdatePlayerRequest struct {
	TeamID      uint   `json:"team_id,string,omitempty"`
	Name        string `json:"name,omitempty" validate:"max=255"`
	Nickname    string `json:"nickname,omitempty" validate:"max=255"`
	Nationality string `json:"nationality,omitempty" validate:"max=255"`
}
