package helpers

import (
	"encoding/json"
	"net/http"
	"quizON/internal/logger"
)

const EmptyResponse = ""

const WrongLoginOrPassword = "Неверный логин или пароль"

type httpError struct {
	Code int
	Err  error
	Body string `json:"error"`
}

func NewHttpError(code int, err error, body string) error {
	return httpError{
		Code: code,
		Err:  err,
		Body: body,
	}
}

func (h httpError) Error() string {
	return h.Error()
}

// HandleHttpError - правильно записывает http ошибки
func HandleHttpError(w http.ResponseWriter, err error) {
	// маппим ошибку к нашей HttpError
	httpErr, ok := err.(httpError)
	// если не удалось замаппить, то что-то сильно идет не по плану
	if !ok {
		logger.Errorf("can't map error: %v", err)
		ResponseWithJson(w, http.StatusInternalServerError, nil)
		return
	}

	// если код ошибки < 500, отвечаем кодом ошибки и телом ошибки
	if httpErr.Code < 500 {
		logger.Info(httpErr.Err)
		ResponseWithJson(w, httpErr.Code, httpErr.Body)
		return
	}

	// если код ошибки >= 500, отвечаем кодом и логируем ошибку
	logger.Error(httpErr.Err)
	w.WriteHeader(httpErr.Code)
	return
}

func ResponseWithJson(w http.ResponseWriter, code int, body interface{}) {
	w.WriteHeader(code)
	if body == nil {
		return
	}

	err := json.NewEncoder(w).Encode(body)
	if err != nil {
		logger.Errorf("can't write body: %v\n", err)
	}

	return
}
