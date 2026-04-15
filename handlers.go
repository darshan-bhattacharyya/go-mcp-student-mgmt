package main

import (
	"context"
	"fmt"

	"github.com/darshan-bhattacharyya/go-mcp-student-mgmt/database"
	"github.com/darshan-bhattacharyya/go-mcp-student-mgmt/models"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type SchoolTools struct {
	DB *database.SchoolDatabase
}

// New SchoolTools initializes the SchoolTools with a database connection
func NewSchoolTools(db *database.SchoolDatabase) *SchoolTools {
	return &SchoolTools{DB: db}
}

// HandleCreateStudent handles "createMessageStudent" requests for the MCP server.
func (st *SchoolTools) HandleCreateStudent(ctx context.Context, req *mcp.CallToolRequest, input *models.Student) (*mcp.CallToolResult, any, error) {

	// Create the student in the database
	rowsAffected, err := st.DB.CreateStudent(input)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create student: %v", err)
	}
	if rowsAffected == 0 {
		return nil, nil, fmt.Errorf("no student was created")
	}

	resultText := fmt.Sprintf("Student %s added to school successfully with ID: %d", input.FirstName, input.ID)

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: resultText},
		},
	}, nil, nil
}
