package models

type Rating struct {
	ID      int    `json:"id" gorm:"primaryKey"` // Primary key (Rating ID)
	UserID  int    `json:"user_id"`              // Foreign key (User ID)
	BookID  int    `json:"book_id"`              // Foreign key (Book ID)
	Rating  int    `json:"rating"`               // Rating (e.g., 1-5)
	Comment string `json:"comment"`              // Comment or review of the book
}
