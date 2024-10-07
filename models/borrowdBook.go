package models

import "time"

type BorrowedBook struct {
	ID         int       `json:"id" gorm:"primaryKey"` // Primary key (Borrowed Book ID)
	UserID     int       `json:"user_id"`              // Foreign key (User ID)
	BookID     int       `json:"book_id"`              // Foreign key (Book ID)
	BorrowedAt time.Time `json:"borrowed_at"`          // Timestamp of when the book was borrowed
	DueDate    time.Time `json:"due_date"`             // Timestamp of the due date for return
}
