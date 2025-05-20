package user

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"review-system-go/internal/db"
)

func ListUsers() ([]User, error) {
	var users []User
	err := db.DB.Find(&users).Error
	return users, err
}

func GetUser(id uint) (*User, error) {
	var user User
	err := db.DB.First(&user, id).Error
	return &user, err
}

func CreateUser(u *User) error {
	return db.DB.Create(u).Error
}

func UpdateUser(u *User) error {
	return db.DB.Save(u).Error
}

func DeleteUser(id uint) error {
	return db.DB.Delete(&User{}, id).Error
}

func Register(username, password, phone string) (*User, error) {
	var count int64
	db.DB.Model(&User{}).Where("name = ?", username).Count(&count)
	if count > 0 {
		return nil, errors.New("username already exists")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	u := &User{
		Name:     username,
		Password: string(hash),
		Phone:    phone,
	}

	return u, db.DB.Create(u).Error
}

func GetByUsername(username string) (*User, error) {
	var u User
	err := db.DB.Where("username = ?", username).First(&u).Error
	if err != nil {
		return nil, err
	}
	return &u, nil
}
