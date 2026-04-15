package models

type Student struct {
	ID            uint          `gorm:"primaryKey" json:"id"`
	FirstName     string        `json:"first_name"`
	LastName      string        `json:"last_name"`
	Email         string        `json:"email"`
	LegalGuardian LegalGuardian `gorm:"foreignKey:LegalGuardianID" json:"legal_guardian"`
}

type LegalGuardian struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}
