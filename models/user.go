package models

import "time"

type Role string

const (
	RoleAdmin     Role = "ADMIN"
	RoleOrganizer Role = "ORGANIZER"
	RoleStaff     Role = "STAFF"
)

type User struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	Role      Role      `json:"role"`
	CreatedAt time.Time `json:"Createdat"`
	UpdatedAt time.Time `json:"UpdatedAt"`
}
