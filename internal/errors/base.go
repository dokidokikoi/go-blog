package errors

import "net/http"

type APPError struct {
	Code       int    `json:"code"`
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

func Success(message string) *APPError {
	return &APPError{
		StatusCode: http.StatusOK,
		Message:    message,
	}
}

func ServerFailed(message string, code int) *APPError {
	return &APPError{
		StatusCode: http.StatusInternalServerError,
		Message:    message,
		Code:       code,
	}
}

func ClientFailed(message string, code int) *APPError {
	return &APPError{
		StatusCode: http.StatusBadRequest,
		Message:    message,
		Code:       code,
	}
}

func ClientAuthnFailed(message string) *APPError {
	return &APPError{
		StatusCode: http.StatusUnauthorized,
		Message:    message,
	}

}
