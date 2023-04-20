package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Socialmedia struct {
	ID             uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name           string `gorm:"not null;" json:"name" form:"name" valid:"required~name is required"`
	SocialMediaURL string `gorm:"not null;" json:"social_media_url" form:"social_media_url" valid:"required~social_media_url is required"`
	UserID         uint
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func (sm *Socialmedia) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(sm)
	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (sm *Socialmedia) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(sm)
	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
