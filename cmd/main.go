package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	// pq is the library that allows Go to interact with PostgreSQL
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432 
	dbname   = "task_manager" 
)

func main() {
	// Retrieve the username from the environment variable
	user := os.Getenv("POSTGRES_USER")
	if user == "" {
		log.Fatal("The POSTGRES_USER environment variable is not set.")
	}

	// Construct the connection string
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable", host, port, user, dbname)

	// Open a connection to the database
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Error opening connection: %s", err)
	}
	defer db.Close()

	// Test the connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error pinging database: %s", err)
	}

	fmt.Println("Successfully connected to database!")
}