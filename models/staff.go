package models

type Staff struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type Teacher struct {
	Staff
	DeptartmentID string `json:"department"`
}

type NonTeachingStaff struct {
	Staff
	Role string `json:"role"`
}
