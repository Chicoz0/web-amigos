package models

import (
	"time"

	"github.com/google/uuid"
)

type Course struct {
	ID                 uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	CreatorID          uuid.UUID `gorm:"type:uuid;not null"`
	Creator            User      `gorm:"foreignKey:CreatorID"`
	Title              string    `gorm:"type:varchar(255);not null"`
	Description        string    `gorm:"type:text"`
	ThumbnailURL       string    `gorm:"type:varchar(255)"`
	GatewayProductID   string    `gorm:"type:varchar(255)"`
	Price              float64   `gorm:"type:decimal(10,2)"`
	AccessDurationDays *int
	IsPublished        bool      `gorm:"default:false"`
	CreatedAt          time.Time `gorm:"autoCreateTime"`

	Modules     []Module
	Lessons     []Lesson
	Enrollments []Enrollment
	Reviews     []CourseReview
}

type Module struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	CourseID  uuid.UUID `gorm:"type:uuid;not null"`
	Course    Course    `gorm:"foreignKey:CourseID"`
	Title     string    `gorm:"type:varchar(255);not null"`
	Position  int       `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`

	Lessons []Lesson
}

type Lesson struct {
	ID              uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	ModuleID        uuid.UUID `gorm:"type:uuid;not null"`
	Module          Module    `gorm:"foreignKey:ModuleID"`
	CourseID        uuid.UUID `gorm:"type:uuid;not null"`
	Course          Course    `gorm:"foreignKey:CourseID"`
	Title           string    `gorm:"type:varchar(255);not null"`
	Description     string    `gorm:"type:text"`
	YoutubeID       string    `gorm:"type:varchar(255)"`
	DurationMinutes int
	Position        int       `gorm:"not null"`
	IsFree          bool      `gorm:"default:false"`
	CreatedAt       time.Time `gorm:"autoCreateTime"`
}
