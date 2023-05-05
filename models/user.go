package models

import "time"

type User struct {
	ID         int64     `json:"id"`
	NAME       string    `json:"name"`
	PHONE      string    `json:"phone"`
	EMAIL      string    `json:"email"`
	PASSWORD   string    `json:"password"`
	CREATED_AT time.Time `json:"created_at"`
}
