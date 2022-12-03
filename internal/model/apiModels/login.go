package apiModels

type LoginRequest struct {
	Login    string `db:"login" json:"login"`
	Password string `db:"password" json:"password"`
}
