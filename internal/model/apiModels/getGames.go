package apiModels

// GetGamesRequest - запрос на получение игр
type GetGamesRequest struct {
	// Page - страница пагинации (начиная с 1)
	Page int32 `json:"page" validate:"required,gte=0,lte=30"`
	// PerPage - количество элементов на странице
	PerPage int32 `json:"per_page" validate:"required,gte=0,lte=30"`
}

// GetGamesResponse - ответ на запрос на получение игр
type GetGamesResponse struct {
	Games []Game `json:"games"`
}
