package model

import "time"

type TeamResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Alias     string    `json:"alias"`
	Origin    string    `json:"origin"`
	Region    string    `json:"region"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateTeamRequest struct {
	Name      string    `json:"name" validate:"required,max=255"`
	Alias     string    `json:"alias" validate:"required,max=3"`
	Origin    string    `json:"origin" validate:"required,max=255"`
	Region    string    `json:"region" validate:"required,max=255"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UpdateTeamRequest struct {
	Name   string `json:"name,omitempty" validate:"max=255"`
	Alias  string `json:"alias,omitempty" validate:"max=3"`
	Origin string `json:"origin,omitempty" validate:"max=255"`
	Region string `json:"region,omitempty" validate:"max=255"`
}
