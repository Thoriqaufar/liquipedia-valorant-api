package model

type TeamDetailsResponse struct {
	ID     uint                     `json:"id"`
	Name   string                   `json:"name"`
	Alias  string                   `json:"alias"`
	Origin string                   `json:"origin"`
	Region string                   `json:"region"`
	Player []ForTeamDetailsResponse `json:"player"`
}

type ForTeamDetailsResponse struct {
	ID          uint   `json:"id"`
	TeamID      uint   `json:"team_id"`
	Name        string `json:"name"`
	Nickname    string `json:"nickname"`
	Nationality string `json:"nationality"`
}

type PlayerDetailsResponse struct {
	ID          uint                     `json:"id"`
	TeamID      uint                     `json:"team_id"`
	Name        string                   `json:"name"`
	Nickname    string                   `json:"nickname"`
	Nationality string                   `json:"nationality"`
	Team        ForPlayerDetailsResponse `json:"team"`
}

type ForPlayerDetailsResponse struct {
	Name string `json:"name"`
}
