package controllers

import (
	"net/http"

	"github.com/abdur-rahman41/hall-management-backend/pkg/domain"
	"github.com/abdur-rahman41/hall-management-backend/pkg/serializer"
	"github.com/labstack/echo/v4"
)

// IAuthController is an interface that defines the methods implemented by the AuthController struct.
type IAuthController interface {
	Signup(e echo.Context) error
	Login(e echo.Context) error
}

// AuthController defines the methods of the IAuthController interface.
type AuthController struct {
	authSvc domain.IAuthService
}

// NewAuthController is a function that returns a new instance of the AuthController struct.
func NewAuthController(authSvc domain.IAuthService) AuthController {
	return AuthController{
		authSvc: authSvc,
	}
}

func (authController *AuthController) Signup(context echo.Context) error {

	registerRequest := &serializer.SignupRequest{}
	// name := context.FormValue("name")
	// studentId := context.FormValue("student_id")
	// email := context.FormValue("email")
	// graduationYear := context.FormValue("graduation_year")
	// role := context.FormValue("role")
	// password := context.FormValue("password")
	// confirmPassword := context.FormValue("confirm_password")

	// // bind the request body to the SignupRequest struct
	// registerRequest := &serializer.SignupRequest{
	// 	ID:        registerRequest.ID,
	// 	Name:      registerRequest.Name,
	// 	RegNumber: registerRequest.RegNumber,
	// 	Email:     registerRequest.Email,
	// 	Phone:     registerRequest.Phone,
	// 	AttachNo:  registerRequest.AttachNo,
	// 	Role:      registerRequest.Role,
	// 	Password:  registerRequest.Password,
	// }

	if err := context.Bind(registerRequest); err != nil {
		return context.JSON(http.StatusBadRequest, "invalid request body")
	}

	// validate the request body
	if err := registerRequest.Validate(); err != nil {
		return context.JSON(http.StatusBadRequest, err.Error())
	}

	// pass the request to the service layer
	if err := authController.authSvc.SignupUser(registerRequest); err != nil {
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusCreated, "user was created successfully")
}

// func (authController *AuthController) Login(e echo.Context) error {
// 	loginRequest := &serializer.LoginRequest{}
// 	if err := e.Bind(loginRequest); err != nil {
// 		return e.JSON(http.StatusBadRequest, "invalid request body")
// 	}

// 	if err := loginRequest.Validate(); err != nil {
// 		return e.JSON(http.StatusBadRequest, err.Error())
// 	}

// 	loginResponse, err := authController.authSvc.Login(loginRequest)

// 	if err != nil {
// 		return e.JSON(http.StatusInternalServerError, err.Error())
// 	}
// 	return e.JSON(http.StatusOK, loginResponse)
// }
