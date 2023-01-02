package apiModels

type RegisterRequest struct {
	GameID      int32  `json:"game_id" validate:"required"`
	TeamName    string `json:"team_name" validate:"required"`
	CaptainName string `json:"captain_name" validate:"required"`
	Phone       string `json:"phone" validate:"required"`
	Telega      string `json:"telega" validate:"required"`
	Amount      int32  `json:"amount" validate:"required,gte=0,lte=8"`
}

type RegisterResponse struct {
	RegistrationNumber int32 `json:"registration_number"`
}
