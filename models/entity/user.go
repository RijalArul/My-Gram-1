package entity

import (
	"my-gram-1/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	Username string `gorm:"not null;unique" form:"username" json:"username" valid:"required~Your username is required"`
	Email    string `gorm:"not null;unique" form:"email" json:"email" valid:"required~Your email is required,email~Invalid email format"`
	Password string `gorm:"not null" form:"password" json:"password" valid:"required~Your password is required,minstringlength(6)~Password has to have a minimum length of 6 characters"`
	Age      int    `gorm:"not null" form:"age" json:"age" valid:"required~Your age is required,range(8|60)~Age minimum Age 8-60"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = helpers.HashPass(u.Password)

	err = nil
	return
}

// func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
// 	_, errUpdate := govalidator.ValidateStruct(u)

// 	if errUpdate != nil {
// 		err = errUpdate
// 		return
// 	}

// 	err = nil
// 	return
// }
