package models

import (
	"errors"
	"mygram/helpers"
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

// User represents the model for users
type User struct {
	ID           uint   `gorm:"primaryKey" json:"id"`
	Username     string `gorm:"not null;uniqueIndex" json:"username" form:"username" valid:"required~Username is required"`
	Email        string `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~Email is required,email~Invalid format"`
	Password     string `gorm:"not null" json:"password" form:"password" valid:"required~Password is required,minstringlength(6)~Password has to be a minimum length of six cahracters"`
	Age          uint   `gorm:"not null;unique" json:"age" form:"age" valid:"required~Age is required"`
	Photos       []Photo
	Comments     []Comment
	Socialmedias []Socialmedia
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)
	if errCreate != nil {
		err = errCreate
		return
	}
	if u.Age < 8 {
		err = errors.New("minimum age is 8")
		return
	}
	u.Password = helpers.HashPass(u.Password)
	err = nil
	return
}
