package model

import "time"

type User struct {
	ID           int        `json:"id"`
	Username     string     `json:"username"`
	Password     string     `json:"password"`
	PasswordHash string     `json:"-" db:"password_hash"`
	FirstName    string     `json:"first_name" db:"first_name"`
	LastName     string     `json:"last_name" db:"last_name"`
	Gender       int        `json:"gender" db:"gender"`
	DateOfBirth  time.Time  `json:"date_of_birth" db:"date_of_birth"`
	Email        string     `json:"email" db:"email"`
	PhoneNumber  string     `json:"phone_number" db:"phone_number"`
	UserType     string     `json:"user_type" db:"user_type"`
	PartnerID    int        `json:"partner_id" db:"partner_id"`
	LastLogin    *time.Time `json:"last_login" db:"last_login"`
	CreatedAt    time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at" db:"updated_at"`
}

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
