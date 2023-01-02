package apiModels

import "time"

type Game struct {
	ID              int32     `json:"id"`
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	Date            time.Time `json:"date"`
	TeamsAmount     int32     `json:"teams_amount"`
	RegisteredTeams int32     `json:"registered_teams"`
	PricePerPerson  int32     `json:"price_per_person"`
	Location        string    `json:"location"`
}

type Registration struct {
	TeamName           string    `json:"team_name"`
	CaptainName        string    `json:"captain_name"`
	Phone              string    `json:"phone"`
	Telega             string    `json:"telega"`
	Amount             int32     `json:"amount"`
	RegistrationNumber int32     `json:"registration_number"`
	CreatedAt          time.Time `json:"created_at"`
}
