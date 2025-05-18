package user

import "time"

type User struct {
	ID             int64     `json:"id"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	Phone          string    `json:"phone"`
	Password       string    `json:"password"`
	CreatedAt      time.Time `json:"created_at"`
	AttentionCount int64     `json:"attention_count"`
	FollowerCount  int64     `json:"follower_count"`
}
