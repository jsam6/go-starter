package user

import (
	"gorm.io/gorm"
	"github.com/jsam6/go-orders-api/config"
)

type User struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email"`
}

var db *gorm.DB

func init() {
	db = config.Connect()
	db.AutoMigrate(&User{})
}

func GetAllMembers() []User {
	var Members []User
	db.Find(&Members)
	return Members
}