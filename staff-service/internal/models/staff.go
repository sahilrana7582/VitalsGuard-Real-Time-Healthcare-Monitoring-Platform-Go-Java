package models

import "time"

type Doctor struct {
	ID                string    `json:"id"`
	TenantID          string    `json:"tenant_id"`
	StaffID           string    `json:"staff_id"`
	Specialization    string    `json:"specialization"`
	LicenseNumber     string    `json:"license_number"`
	YearsOfExperience int       `json:"years_of_experience"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

type Nurse struct {
	ID            string    `json:"id"`
	TenantID      string    `json:"tenant_id"`
	StaffID       string    `json:"staff_id"`
	Shift         string    `json:"shift"`
	FloorAssigned string    `json:"floor_assigned"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type Staff struct {
	ID            string     `db:"id"`
	TenantID      string     `db:"tenant_id"`
	UserID        string     `db:"user_id"`
	FullName      string     `db:"full_name"`
	Gender        string     `db:"gender"`
	DOB           *time.Time `db:"dob"`
	ContactNumber string     `db:"contact_number"`
	Email         string     `db:"email"`
	Address       string     `db:"address"`
	Department    string     `db:"department"`
	CreatedAt     time.Time  `db:"created_at"`
	UpdatedAt     time.Time  `db:"updated_at"`
}
