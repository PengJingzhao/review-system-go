package user

import "review-system-go/internal/db"

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
