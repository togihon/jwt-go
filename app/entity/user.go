package entity

import (
	"jwt-go/pkg/helper"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	FullName string    `gorm:"not null" json:"full_name" form:"full_name" valid:"required~Your full name is required"`
	Email    string    `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~Your email is required,email~Invalid email format"`
	Password string    `gorm:"not null" json:"password" form:"password" valid:"required~Your password is required,minstringlength(6)~Password must be 6 characters or more"`
	Level    string    `gorm:"not null" json:"level" form:"level" valid:"required~Level is required"`
	Products []Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"products"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = helper.HashPass(u.Password)
	return nil
}
