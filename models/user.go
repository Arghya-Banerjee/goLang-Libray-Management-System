package models

type User struct {
	ID       int    `json:"id" gorm:"primaryKey"` // Primary key (User ID)
	Name     string `json:"name"`                 // Name of the user
	Email    string `json:"email" gorm:"unique"`  // Email (unique identifier for each user)
	Password string `json:"-"`                    // User's password (hidden in responses)
}
