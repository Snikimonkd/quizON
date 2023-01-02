package apiModels

import "time"

type CreateGameRequest struct {
	Name           string    `json:"name" validate:"required"`
	Description    string    `json:"description"  validate:"required"`
	Date           time.Time `json:"date"  validate:"required"`
	TeamsAmount    int32     `json:"teams_amount"  validate:"required"`
	PricePerPerson int32     `json:"price_per_person"  validate:"required"`
	Location       string    `json:"location"  validate:"required"`
}

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

type CreateGameResponse struct {
	Game
	//ID              int32     `json:"id"`
	//Name            string    `json:"name"`
	//Description     string    `json:"description"`
	//Date            time.Time `json:"date"`
	//TeamsAmount     int32     `json:"teams_amount"`
	//RegisteredTeams int32     `json:"registered_teams"`
	//PricePerPerson  int32     `json:"price_per_person"`
	//Location        string    `json:"location"`
}
