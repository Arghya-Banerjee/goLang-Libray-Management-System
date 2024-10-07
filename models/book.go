package models

type Book struct {
	ID     int     `json:"id" gorm:"primaryKey"` // Primary key (Book ID)
	Title  string  `json:"title"`                // Title of the book
	Author string  `json:"author"`               // Author of the book
	Genre  string  `json:"genre"`                // Genre of the book
	Stock  int     `json:"stock"`                // Stock available for the book
	Rating float64 `json:"rating"`               // Average rating of the book
}
