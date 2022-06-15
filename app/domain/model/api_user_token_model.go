package model

import "time"

// User struct
type ApiUserToken struct {
	ID        int       `json:"id" gorm:"primary_key"`
	User_ID   int       `json:"user_id" gorm:"type:int;not null"`
	Token     string    `json:"token" gorm:"type:text;not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
