package models

import (
	"errors"
	"html"
	"strings"

	"github.com/jinzhu/gorm"
)

//var server = controllers.Server{}

type UserDetails struct {
	Email    string `gorm:"size:255;not null;unique" json:"email"`
	Password string `gorm:"size:100;not null;unique" json:"password"`
	Number   string `gorm:"size:100;not null;unique" json:"number"`
}

//
func (u *UserDetails) Prepare() {
	u.Password = html.EscapeString(strings.TrimSpace(u.Password))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.Number = html.EscapeString(strings.TrimSpace(u.Number))
}

//
func (u *UserDetails) FindUserByEmailID(email interface{}) (*UserDetails, error) {
	var err error
	//err = controllers.Global.DB.Debug().Model(User{}).Where("email like ?", email).Take(&u).Error
	if err != nil {
		return &UserDetails{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &UserDetails{}, errors.New("User not found")
	}
	return u, err

}
