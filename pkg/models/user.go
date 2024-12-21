package model

import (
	"fmt"
	"strconv"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID             string `json:"id" gorm:"uniqueIndex;size:7"`
	Name           string `json:"name"`
	RegNumber      string `json:"reg_number" gorm:"uniqueIndex;size:5"`
	Email          string `json:"email" gorm:"uniqueIndex"`
	Phone          string `json:"phone"`
	AttachNo       string `json:"attach_no" gorm:"uniqueIndex"`
	Role           string `json:"role"`
	Password       string `json:"password"` // Password is never exposed
	Session        string `json:"session"`
	IsUserVerified bool   `json:"is_user_verified"`
}

func (x *User) SetVerificationProperties() {
	x.IsUserVerified = true
	x.Session = generateSession(x.ID)
}

func generateSession(StudentId string) string {
	FirstNumber := StudentId[0:2]
	FirstNumberInt, err := strconv.Atoi(FirstNumber)
	if err != nil {
		return ""
	}
	SecondNumber := strconv.Itoa(FirstNumberInt - 1)
	return fmt.Sprintf("20%s-%s", SecondNumber, FirstNumber)
}
