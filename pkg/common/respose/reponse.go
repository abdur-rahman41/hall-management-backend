package response

import (
	"errors"
	"fmt"
	"net/http"
)

type body map[string]interface{}

type SuccessResponse struct {
	Message string
	Data    interface{}
}

// Map for errors with http code
var ResponseCode = make(map[string]int, 0)

var (
	ErrParsingRequestBody   = NewError("failed to parse request body", http.StatusBadRequest)
	ErrInvalidRequestParams = NewError("request params is not valid", http.StatusBadRequest)
	ErrEventAlreadyExists   = NewError("Event already exists", http.StatusBadRequest)
)

func responseMap() map[string]int {
	return ResponseCode
}

func NewError(message string, httpCode int) error {
	_, available := ResponseCode[message]
	if !available {
		ResponseCode[message] = httpCode
	}
	return errors.New(message)
}

type StudentIDExistsError struct {
	ID string
}

func (e *StudentIDExistsError) Error() string {
	return fmt.Sprintf("Student ID %s already exists", e.ID)
}

type EmailExistsError struct {
	Email string
}

func (e *EmailExistsError) Error() string {
	fmt.Println(e.Email)
	return fmt.Sprintf("Email %s already exists", e.Email)
}

type UserNotVerifiedError struct{}

func (e *UserNotVerifiedError) Error() string {
	return "Please wait until verified"
}

func GenerateErrorResponseBody(err error) (int, body) {
	message := err.Error()
	return readFromMap(message)
}

func readFromMap(message string) (int, body) {
	httpStatus, available := responseMap()[message]
	if available {
		return httpStatus, generateResponseBody(message)
	}
	return http.StatusInternalServerError, generateResponseBody(message)
}

func generateResponseBody(message string) body {
	return body{
		"message": message,
	}
}

func GenerateSuccessResponse(message string, data interface{}) *SuccessResponse {
	return &SuccessResponse{
		Message: message,
		Data:    data,
	}

}
