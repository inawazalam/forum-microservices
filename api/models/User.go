package models

import (
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"github.com/jinzhu/gorm"

	"golang.org/x/crypto/bcrypt"
)

//
type UserDetail struct {
	autherID  uint64
	nickname  string
	email     string
	picurl    string
	vehicleID string
}

var autherID uint64
var nickname string
var userEmail string
var picurl string
var vehicleID string

//User model
type User struct {
	Nickname  string `gorm:"size:255;not null;unique" json:"nickname"`
	Email     string `gorm:"size:100;not null;unique" json:"email"`
	VehicleID string `gorm:"size:100;not null;unique" json:"vehicleid"`
	Picurl    string `gorm:"size:100;not null;unique" json:"profile_pic_url"`
	//Password  string    `gorm:"size:100;not null;" json:"password"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
}

//
/*func Prepare() {
	var u User
	u.Nickname = nickname
	u.Email = email
	u.VehicleID = vehicleID
	u.CreatedAt = time.Now()
	u.Picurl = picurl
}*/

//Hash for password
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

//VerifyPassword compare password and hashcode
func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

/*func (u *User) BeforeSave() error {
	hashedPassword, err := Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}*/

//Prepare initilize user object

//Validate User
func (u *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.Nickname == "" {
			return errors.New("Required Nickname")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil

	case "login":
		if u.Nickname == "" {
			return errors.New("Required Nickname")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil
	default:
		if u.Nickname == "" {
			return errors.New("Required Nickname")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil
	}
}

//
func FindUserByEmail(email string, db *gorm.DB) (*uint64, error) {
	var err error
	var id uint64
	var number *uint64
	var name string
	var profile_pic_url string
	var uuid string
	userEmail = email
	//var

	row := db.Table("user_login").Where("email LIKE ?", email).Select("id,number").Row()

	row.Scan(&id, &number)
	autherID = id
	row1 := db.Table("user_details").Where("user_id = ?", id).Select("name, profile_pic_url").Row()
	row1.Scan(&name, &profile_pic_url)
	nickname = name
	picurl = profile_pic_url

	row2 := db.Table("vehicle_details").Where("owner_id = ?", id).Select("uuid").Row()
	row2.Scan(&uuid)
	vehicleID = uuid
	return number, err
}
