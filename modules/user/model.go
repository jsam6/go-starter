package user

import (
	"github.com/jsam6/go-orders-api/config"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name  string
	Email string
}

var db *gorm.DB

func init() {
	db = config.Connect()
	db.AutoMigrate(&User{})
}

func GetAllUsers() []User {
	var Users []User
	db.Find(&Users)
	return Users
}

func GetUserById(id string) (*User, *gorm.DB) {
	var getMember User
	db := db.Where("id=?", id).Find(&getMember)
	return &getMember, db
}

func (m *User) CreateUser() *User {
	db.Save(m)
	db.Create(&m)
	return m
}