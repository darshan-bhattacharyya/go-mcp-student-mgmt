package database

import (
	"fmt"
	"os"

	"github.com/darshan-bhattacharyya/go-mcp-student-mgmt/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type SchoolDatabase struct {
	DB *gorm.DB
}

// NewSchoolDatabase initializes the database connection using environment variables and returns a SchoolDatabase instance
func NewSchoolDatabase() (*SchoolDatabase, error) {
	host := os.Getenv("PG_HOST")
	user := os.Getenv("PG_USER")
	password := os.Getenv("PG_PASSWORD")
	dbname := os.Getenv("PG_DBNAME")
	port := os.Getenv("PG_PORT")
	sslmode := os.Getenv("PG_SSLMODE")
	timezone := os.Getenv("PG_TIMEZONE")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		host, user, password, dbname, port, sslmode, timezone)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	// Auto-migrate models
	err = db.AutoMigrate(&models.Student{}, &models.LegalGuardian{}, &models.Staff{}, &models.Teacher{}, &models.NonTeachingStaff{})
	if err != nil {
		return nil, err
	}
	return &SchoolDatabase{DB: db}, nil
}

// Student CRUD methods
func (db *SchoolDatabase) CreateStudent(student *models.Student) (int64, error) {
	result := db.DB.Create(student)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}

func (db *SchoolDatabase) GetStudentByID(id uint) (*models.Student, error) {
	var student models.Student
	err := db.DB.Preload("LegalGuardian").First(&student, id).Error
	return &student, err
}

func (db *SchoolDatabase) UpdateStudent(student *models.Student) (int64, error) {
	result := db.DB.Save(student)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}

func (db *SchoolDatabase) DeleteStudent(id uint) error {
	return db.DB.Delete(&models.Student{}, id).Error
}

func (db *SchoolDatabase) GetAllStudents() ([]models.Student, error) {
	var students []models.Student
	err := db.DB.Preload("LegalGuardian").Find(&students).Error
	return students, err
}

// Staff CRUD methods
func (db *SchoolDatabase) CreateStaff(staff *models.Staff) (int64, error) {
	result := db.DB.Create(staff)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}

func (db *SchoolDatabase) GetStaffByID(id uint) (*models.Staff, error) {
	var staff models.Staff
	err := db.DB.First(&staff, id).Error
	return &staff, err
}

func (db *SchoolDatabase) UpdateStaff(staff *models.Staff) (int64, error) {
	result := db.DB.Save(staff)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}

func (db *SchoolDatabase) DeleteStaff(id uint) (int64, error) {
	result := db.DB.Delete(&models.Staff{}, id)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}
