package session

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/andrew-hayworth22/wodify-go"
)

// ErrorType classifies the type of error that occurred.
type ErrorType string

const (
	ErrUnauthorized ErrorType = "unauthorized"
	ErrForbidden    ErrorType = "forbidden"
	ErrValidation   ErrorType = "validation"
	ErrNotFound     ErrorType = "not_found"
	ErrRateLimited  ErrorType = "rate_limited"
	ErrNetwork      ErrorType = "network"
	ErrUnknown      ErrorType = "unknown"
)

// Error represents an error that occurred during the execution of an API request.
type Error struct {
	Type    ErrorType `json:"type"`
	Message string    `json:"message"`
	Detail  string    `json:"detail,omitempty"`
	Status  int       `json:"status,omitempty"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s: %s", e.Type, e.Message)
}

// FormatError formats an error into an Error type for use in Wails error formatter.
func FormatError(err error) any {
	return toError(err)
}

func toError(err error) *Error {
	if err == nil {
		return nil
	}

	var appErr *Error
	if errors.As(err, &appErr) {
		return appErr
	}

	var apiErr *wodify.APIError
	if errors.As(err, &apiErr) {
		e := &Error{
			Message: apiErr.UserMessage,
			Detail:  apiErr.DeveloperMessage,
			Status:  apiErr.HTTPCode,
		}
		switch e.Status {
		case http.StatusUnauthorized:
			e.Type = ErrUnauthorized
		case http.StatusForbidden:
			e.Type = ErrForbidden
		case http.StatusBadRequest, http.StatusUnprocessableEntity:
			e.Type = ErrValidation
		case http.StatusNotFound:
			e.Type = ErrNotFound
		case http.StatusTooManyRequests:
			e.Type = ErrRateLimited
		default:
			e.Type = ErrUnknown
		}
		if e.Message == "" {
			e.Message = defaultMessage(e.Type)
		}
	}

	return &Error{
		Type:    ErrUnknown,
		Message: err.Error(),
	}
}

func defaultMessage(errType ErrorType) string {
	switch errType {
	case ErrUnauthorized:
		return "Your API key is no longer valid. Please enter it again."
	case ErrForbidden:
		return "Your API key doesn't have access to this resource."
	case ErrValidation:
		return "Invalid request"
	case ErrNotFound:
		return "That record couldn't be found."
	case ErrRateLimited:
		return "Too many requests. Please wait a moment and try again."
	default:
		return "Something went wrong."
	}
}
