package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Comment struct {
	ID        uint `gorm:"primaryKey" json:"id"`
	UserID    uint
	PhotoID   uint
	Message   string `gorm:"not null;" json:"message" form:"message" valid:"required~message is required"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (cmnt *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(cmnt)
	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (cmnt *Comment) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(cmnt)
	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
