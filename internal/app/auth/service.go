package auth

import "review-system-go/internal/app/user"

func RegisterUser(username, password, phone string) error {
	_, err := user.Register(username, password, phone)
	if err != nil {
		return err
	}
	return nil
}
