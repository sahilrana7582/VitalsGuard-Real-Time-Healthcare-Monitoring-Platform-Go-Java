package models

import "time"

type Patient struct {
	ID              string    `db:"id"`
	TenantID        string    `db:"tenant_id"`
	FullName        string    `db:"full_name"`
	Age             int       `db:"age"`
	AdmissionReason string    `db:"admission_reason"`
	CreatedAt       time.Time `db:"created_at"`
	UpdatedAt       time.Time `db:"updated_at"`
}

type PatientProfile struct {
	ID            string    `db:"id"`
	TenantID      string    `db:"tenant_id"`
	PatientID     string    `db:"patient_id"`
	Gender        string    `db:"gender"`
	DOB           time.Time `db:"dob"`
	BloodGroup    string    `db:"blood_group"`
	ContactNumber string    `db:"contact_number"`
	Email         string    `db:"email"`
	Address       string    `db:"address"`
	PostalCode    string    `db:"postal_code"`
	CreatedAt     time.Time `db:"created_at"`
	UpdatedAt     time.Time `db:"updated_at"`
}

type VitalsInfo struct {
	ID                 string    `db:"id"`
	TenantID           string    `db:"tenant_id"`
	PatientID          string    `db:"patient_id"`
	HeartRate          int       `db:"heart_rate"`
	BloodPressureUpper *int      `db:"blood_pressure_upper"`
	BloodPressureLower *int      `db:"blood_pressure_lower"`
	RespiratoryRate    *int      `db:"respiratory_rate"`
	OxygenSaturation   *int      `db:"oxygen_saturation"`
	Temperature        *float32  `db:"temperature"`
	GlucoseLevel       *float32  `db:"glucose_level"`
	RecordedAt         time.Time `db:"recorded_at"`
	CreatedAt          time.Time `db:"created_at"`
	UpdatedAt          time.Time `db:"updated_at"`
}
