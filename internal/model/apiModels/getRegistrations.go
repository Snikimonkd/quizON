package apiModels

type GetRegistrationsRequest struct {
	GameID int32 `json:"game_id"`
}

type GetRegistrationsResponse struct {
	Registrations []Registration `json:"registrations"`
}
