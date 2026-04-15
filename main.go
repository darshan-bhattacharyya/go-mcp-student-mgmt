package main

import (
	"log"

	"github.com/darshan-bhattacharyya/go-mcp-student-mgmt/database"
)

func main() {
	db, err := database.NewSchoolDatabase()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
}
