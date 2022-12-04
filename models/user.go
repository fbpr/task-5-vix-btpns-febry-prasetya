package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/fbpr/task-5-vix-btpns-febry-prasetya/helpers"
	"gorm.io/gorm"
)

type User struct {
	ID uint `gorm:"not null;primaryKey" json:"id"`
	Username string `gorm:"not null;type:varchar(255)" valid:"required" json:"username"`
	Email string `gorm:"not null;unique;type:varchar(255)" valid:"required,email" json:"email"`
	Password string `gorm:"not null;unique;type:varchar(255)" valid:"required,minstringlength(6)" json:"password"`
	Photos []Photo `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) { 
	_, err = govalidator.ValidateStruct(u) 
	if err != nil {
		return 
	} 
	hashedPassword, _ := helpers.Hash(u.Password)
	u.Password = string(hashedPassword)
	
	return nil
} 

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) { 
	_, err = govalidator.ValidateStruct(u) 
	if err != nil {
		return 
	}
	hashedPassword, _ := helpers.Hash(u.Password)
	u.Password = string(hashedPassword) 
	
	return nil
}