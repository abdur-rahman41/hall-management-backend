package serializer

import (
	"errors"

	validation "github.com/go-ozzo/ozzo-validation"
)

// SignupRequest defines the request body for the signup request.
type SignupRequest struct {
	ID        string `json:"id" gorm:"uniqueIndex"`
	Name      string `json:"name"`
	RegNumber string `json:"reg_number"`
	Email     string `json:"email" gorm:"uniqueIndex"`
	Phone     string `json:"phone"`
	AttachNo  string `json:"attach_no"`
	Role      string `json:"role"`
	Password  string `json:"password"` // Password is never exposed
}

type LoginRequest struct {
	ID       *string `json:"id,omitempty"`
	Email    *string `json:"email,omitempty"`
	Password string  `json:"password"`
}

type LoginResponse struct {
	Name           string `json:"name"`
	Email          string `json:"email"`
	StudentId      string `json:"student_id"`
	IsUserVerified bool   `json:"is_user_verified"`
	IsActive       bool   `json:"is_active"`
	Role           string `json:"role"`
	AccessToken    string `json:"access_token"`
}

// Validate validates the request body for the SignupRequest.
func (request SignupRequest) Validate() error {
	return validation.ValidateStruct(&request,
		validation.Field(&request.Name, validation.Required.Error("Name cannot be empty"), validation.Length(4, 128)),
		validation.Field(&request.ID, validation.Required.Error("StudentId must be 7 character"), validation.Length(7, 7)),
		validation.Field(&request.Email, validation.Required.Error("Email cannot be empty"), validation.Length(4, 128)),
		validation.Field(&request.RegNumber, validation.Required.Error("Registration cannot be empty"), validation.Length(5, 5)),
		validation.Field(&request.Role, validation.Required.Error("Role cannot be empty")),
		validation.Field(&request.AttachNo, validation.Required.Error("Attach Number cannot be empty"), validation.Length(8, 128)),
		validation.Field(&request.Phone, validation.Required.Error("Phone Number cannot be empty"), validation.Length(8, 128)),
		validation.Field(&request.Password, validation.Required.Error("Password cannot be empty"), validation.Length(8, 128)),
	)
}

func (request LoginRequest) Validate() error {
	return validation.ValidateStruct(&request,
		validation.Field(&request.Password, validation.Required.Error("Password cannot be empty"), validation.Length(8, 128)),
		validation.Field(&request.ID, validation.By(func(value interface{}) error {
			if request.ID == nil && request.Email == nil {
				return errors.New("must give StudentId or Email")
			}
			return nil
		})),
	)
}
