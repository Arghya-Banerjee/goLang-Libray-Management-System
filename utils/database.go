package utils

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb" // SQL Server driver
)

func InitDB() *sql.DB {

	server := "localhost"    // Use "localhost" or "127.0.0.1"
	port := 1433             // Default SQL Server port
	user := "sa"             // SQL Server username
	password := "qwaszx1234" // SQL Server password
	database := "library_db" // Database name

	// Connection string for MS SQL Server
	connectionString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		server, user, password, port, database)

	// Open the connection
	db, err := sql.Open("sqlserver", connectionString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err)
	}

	// Test the connection
	err = db.Ping()
	if err != nil {
		log.Fatal("Error pinging database: ", err)
	}

	fmt.Println("Connected to MS SQL Server!")
	return db
}
