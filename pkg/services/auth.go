package services

import (
	"github.com/abdur-rahman41/hall-management-backend/pkg/domain"
	model "github.com/abdur-rahman41/hall-management-backend/pkg/models"
	"github.com/abdur-rahman41/hall-management-backend/pkg/serializer"
	"github.com/abdur-rahman41/hall-management-backend/pkg/utils"
)

// authService defines the methods of the domain.IAuthService interface.
type authService struct {
	authRepo domain.IAuthRepo
}

// AuthServiceInstance returns a new instance of the authService struct.
func AuthServiceInstance(authRepo domain.IAuthRepo) domain.IAuthService {
	return &authService{
		authRepo: authRepo,
	}
}

// SignupUser creates a new user with the given user details.
func (service *authService) SignupUser(registerRequest *serializer.SignupRequest) error {
	// Check if the user already exists
	err := service.authRepo.DuplicateUserChecker(&registerRequest.ID, &registerRequest.Email)
	if err != nil {
		return err
	}

	// get hashed password
	passwordHash, err := utils.GetPasswordHash(registerRequest.Password)
	if err != nil {
		return err
	}

	// create user
	user := &model.User{
		ID:        registerRequest.ID,
		Name:      registerRequest.Name,
		RegNumber: registerRequest.RegNumber,
		Email:     registerRequest.Email,
		Phone:     registerRequest.Phone,
		AttachNo:  registerRequest.AttachNo,
		Role:      registerRequest.Role,
		Password:  passwordHash,
	}

	user.SetVerificationProperties()
	// user.OtpExpiryTime = time.Now()
	// //? implement verification later

	// //Send verification email to user
	// err = email.SendEmail(user.Email, email.UserVerificationSubject, email.UserVerificationTemplate)
	// if err != nil {
	// 	return err
	// }

	//////Notify admin
	// emailBody, err := email.CreateAdminNotificationEmail(user.Name, user.StudentId)
	// if err != nil {
	// 	return err
	// }

	// adminEmail := "ice.alumni.management.system@gmail.com"
	// err = email.SendEmail(adminEmail, email.AdminNotificationSubject, emailBody)
	// if err != nil {
	// 	return err
	// }

	if err := service.authRepo.CreateUser(user); err != nil {
		return err
	}

	return nil
}

func (service *authService) Login(loginRequest *serializer.LoginRequest) (*serializer.LoginResponse, error) {
	// Check user is verified or not
	var identifier *string
	// if studentId or email is not provided it gets an error from the validation in the controller layer
	if loginRequest.Email != nil {
		identifier = loginRequest.Email
	} else {
		identifier = loginRequest.ID
	}

	user, err := service.authRepo.FindAuthorizedUserByEmailOrStudentId(identifier)
	if err != nil {
		return nil, err
	}

	// Check password
	if err := utils.CheckPassword(user.Password, loginRequest.Password); err != nil {
		return nil, err
	}

	// Create JWT token
	accessToken, err := utils.GetJwtForUser(user)
	if err != nil {
		return nil, err
	}

	return &serializer.LoginResponse{
		Name:           user.Name,
		Email:          user.Email,
		StudentId:      user.ID,
		IsUserVerified: user.IsUserVerified,
		IsActive:       true,
		Role:           user.Role,
		AccessToken:    accessToken,
	}, nil
}
