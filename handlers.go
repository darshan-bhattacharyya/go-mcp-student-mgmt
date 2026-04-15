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

// HandleViewStudents handles "view_student" requests to view all students and their guardians
func (st *SchoolTools) HandleViewStudents(ctx context.Context, req *mcp.CallToolRequest, input any) (*mcp.CallToolResult, any, error) {

	// Retrieve all students from the database
	students, err := st.DB.GetAllStudents()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to retrieve students: %v", err)
	}

	// Format the response
	var responseText string
	if len(students) == 0 {
		responseText = "No students found in the database."
	} else {
		responseText = fmt.Sprintf("Found %d student(s) in the database:\n\n", len(students))
		for _, student := range students {
			responseText += fmt.Sprintf("ID: %d\n", student.ID)
			responseText += fmt.Sprintf("Name: %s %s\n", student.FirstName, student.LastName)
			responseText += fmt.Sprintf("Email: %s\n", student.Email)
			if student.LegalGuardian != nil {
				responseText += fmt.Sprintf("Student: %s %s and their Legal Guardian: %s %s (Email: %s)\n", student.FirstName, student.LastName, student.LegalGuardian.FirstName, student.LegalGuardian.LastName, student.LegalGuardian.Email)
			}
			responseText += "---\n"
		}
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: responseText},
		},
	}, nil, nil
}
