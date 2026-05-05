package models

import "time"

type EventType string

const (
	EventTypeSeminar    EventType = "SEMINAR"
	EventTypeWorkshop   EventType = "WORKSHOP"
	EventTypeConference EventType = "CONFERENCE"
)

type Event struct {
	ID            uint              `json:"id"`
	Title         string            `json:"title"`
	Type          string            `json:"type"`
	Location      string            `json:"location"`
	StartTime     time.Time         `json:"startTime"`
	EndTime       time.Time         `json:"EndTime"`
	Capacity      int               `json:"capacity"`
	Description   string            `json:"Description"`
	ImgUrl        string            `json:"img_url"`
	Status        string            `json:"status" gorm:"size:20;default:'pending'"` // pending, approved, rejected
	ReviewedBy    *uint             `json:"reviewedBy"`
	ReviewedAt    *time.Time        `json:"Reviewed_at"`
	// Registrations []EventAttendence `json:"registration" gorm:"foreignKey:EventID;constraint:OnDelete:CASCADE"`
	// // Many-to-Many: Event has many Attendees through EventAttendee
	// Attendees []Attendence `json:"attendence" gorm:"many2many:event_attendees;"`
	CreatedAt time.Time    `json:"CreatedAt"`
	UpdatedAt time.Time    `json:"UpdatedAt"`
}
