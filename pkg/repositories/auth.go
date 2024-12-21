package repositories

import (
	"errors"

	response "github.com/abdur-rahman41/hall-management-backend/pkg/common/respose"
	"github.com/abdur-rahman41/hall-management-backend/pkg/domain"
	model "github.com/abdur-rahman41/hall-management-backend/pkg/models"
	"gorm.io/gorm"
)

// userRepo defines the methods of the domain.IUserRepo interface.
type authRepo struct {
	db *gorm.DB
}

// UserDBInstance returns a new instance of the userRepo struct.
func AuthDBInstance(d *gorm.DB) domain.IAuthRepo {
	return &authRepo{
		db: d,
	}
}

//DuplicateUserChecker returns a user model by the username.

func (repo *authRepo) DuplicateUserChecker(ID *string, Email *string) error {
	user := &model.User{}
	if err := repo.db.Where("id= ?", ID).First(user).Error; err == nil {
		return &response.StudentIDExistsError{ID: *ID}
	}
	if err := repo.db.Where("email = ?", Email).First(user).Error; err == nil {
		return &response.EmailExistsError{Email: *Email}
	}
	return nil
}

// CreateUser creates a new user with given user details.
func (repo *authRepo) CreateUser(user *model.User) error {
	if err := repo.db.Create(user).Error; err != nil {
		//log.Fatal(err)
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return errors.New("user already exists")
		}
		return err
	}
	return nil
}

func (repo *authRepo) FindAuthorizedUserByEmailOrStudentId(value interface{}) (*model.User, error) {
	user := &model.User{}
	if err := repo.db.Where("id = ? OR email = ?", value, value).First(user).Error; err != nil {
		return nil, err
	}

	if !user.IsUserVerified {
		return nil, &response.UserNotVerifiedError{}
	}

	return user, nil
}
