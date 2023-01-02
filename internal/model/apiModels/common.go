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
