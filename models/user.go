package models

type User struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"` // Ensure the email is unique
	Password string `json:"password"`            // The password field
}
