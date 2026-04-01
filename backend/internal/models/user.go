package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Email        string    `gorm:"type:varchar(255);uniqueIndex;not null"`
	PasswordHash string    `gorm:"type:varchar(255);not null"`
	FullName     string    `gorm:"type:varchar(255);not null"`
	Role         string    `gorm:"type:varchar(50);not null"` // Ex: admin, creator, student
	CreatedAt    time.Time `gorm:"autoCreateTime"`

	Courses     []Course       `gorm:"foreignKey:CreatorID"`
	Enrollments []Enrollment   `gorm:"foreignKey:UserID"`
	Reviews     []CourseReview `gorm:"foreignKey:UserID"`
}
