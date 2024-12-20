package domain

import (
	model "github.com/abdur-rahman41/hall-management-backend/pkg/models"
	"github.com/abdur-rahman41/hall-management-backend/pkg/serializer"
)

type IAuthRepo interface {
	DuplicateUserChecker(StudentId *string, Email *string) error
	CreateUser(user *model.User) error
	//FindAuthorizedUserByEmailOrStudentId(interface{}) (*model.User, error)
}

type IAuthService interface {
	SignupUser(registerRequest *serializer.SignupRequest) error
	//Login(loginRequest *serializer.LoginRequest) (*serializer.LoginResponse, error)
}
