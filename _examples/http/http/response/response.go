package response

import (
	"encoding/json"
	"net/http"
)

type HTTPError struct {
	Error Err `json:"error"`
}

type Err struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func Success(w http.ResponseWriter, statusCode int, payload any) error {
	res, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(res)

	return nil
}

func Error(w http.ResponseWriter, statusCode int) error {
	statusText := http.StatusText(statusCode)
	if statusText == "" {
		panic("invalid status code")
	}

	httpError := HTTPError{Error: Err{
		Code:    statusCode,
		Message: statusText,
	}}

	res, err := json.Marshal(httpError)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(res)

	return nil
}
