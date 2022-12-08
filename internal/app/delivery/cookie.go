package delivery

import (
	"net/http"
	"quizON/internal/config"
	"quizON/internal/model/postgres/public/model"
)

const CookieName = "token"

func SetCookie(w http.ResponseWriter, dbCookie model.Cookies) {
	cookie := http.Cookie{
		Name:     CookieName,
		Value:    dbCookie.Value.String(),
		Expires:  dbCookie.ExpiresAt,
		SameSite: http.SameSiteStrictMode,
		Domain:   config.GlobalConfig.Cookie.Domain,
		Secure:   config.GlobalConfig.Cookie.Secure,
		HttpOnly: true,
		Path:     "/",
	}

	http.SetCookie(w, &cookie)
}
