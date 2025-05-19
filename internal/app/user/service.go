package user

import "review-system-go/internal/db"

func ListUsers() ([]User, error) {
	var users []User
	err := db.DB.Find(&users).Error
	return users, err
}
