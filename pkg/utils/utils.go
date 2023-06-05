package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Позволяет парсить json-строку, содержащуюся в запросе
func ParseBody(r *http.Request, x interface{}) error {
	if body, err := io.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return err
		}
	}
	return nil
}

// Позволяет отправить ошибку со статусом 500
func SendHTTPInternalError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(w, "Internal error: %s", err)
}

// Позволяет отправить json со статусом 200,
// включая в себя вывод элемента
func SendHTTPSuccessResult(w http.ResponseWriter, item ...interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", item)
}

// Позволяет отправить ошибку 404
// Функция была создана для обработки 404 статуса в контроллерах
func SendHTTPNotFoundError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "Not found: %s", err)
}

// Позволяет отправить кастомную ошибку с указанием статус-кода
func SendCustomError(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)
	fmt.Fprintf(w, "Error: %s", err)
}
