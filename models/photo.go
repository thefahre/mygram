package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photo struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Title     string `gorm:"not null;" json:"title" form:"title" valid:"required~title is required"`
	Caption   string `json:"caption" form:"caption"`
	PhotoURL  string `gorm:"not null;" json:"photo_url" form:"photo_url" valid:"required~photo_url is required"`
	UserID    uint
	Comments  []Comment
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (p *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)
	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (p *Photo) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)
	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
