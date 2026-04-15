package main

import (
	"log"
	"net/http"

	"github.com/darshan-bhattacharyya/go-mcp-student-mgmt/database"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func runMCPServer(schoolTools *SchoolTools) {
	impl := &mcp.Implementation{
		Name:    "hello-world-server",
		Version: "v1.0.0",
	}
	server := mcp.NewServer(impl, nil)

	mcp.AddTool(server, &mcp.Tool{
		Description: "Add a student to the school database",
		Name:        "add_student",
	}, schoolTools.HandleCreateStudent)

	handler := mcp.NewStreamableHTTPHandler(
		func(r *http.Request) *mcp.Server { return server },
		nil,
	)

	log.Println("Starting MCP Server on http://localhost:8081 ...")
	if err := http.ListenAndServe(":8081", handler); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func main() {
	db, err := database.NewSchoolDatabase()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	schoolTools := NewSchoolTools(db)
	runMCPServer(schoolTools)
}
