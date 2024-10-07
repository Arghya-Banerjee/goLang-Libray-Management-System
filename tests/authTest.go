package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := "password123"
	storedHash := "$2a$10$trrc5T0ZLch9Uj9i.65KNu2dejVp.N.DSYdJD9Mk/KdTGIBZNxXeu"

	// Manually hash the password
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	fmt.Println("Manually hashed password:", string(hashedPassword))

	// Compare the provided password with the stored hash
	err := bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(password))
	if err != nil {
		fmt.Println("Passwords do not match:", err)
	} else {
		fmt.Println("Passwords match!")
	}
}
