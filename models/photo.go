package models

import (
	"gorm.io/gorm"
	"github.com/asaskevich/govalidator"
)

type Photo struct {
	ID uint `gorm:"not null;primaryKey" json:"id"`
	Title string `gorm:"type:varchar(255)" valid:"required" json:"title"`
	Caption string `gorm:"type:varchar(255)" valid:"required" json:"caption"`
	PhotoUrl string `gorm:"type:varchar(255)" valid:"required" json:"photo_url"`
	UserID uint `json:"user_id"`
	User *User
}

func (p *Photo) BeforeCreate(tx *gorm.DB) (err error) { 
	_, err = govalidator.ValidateStruct(p) 
	if err != nil {
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