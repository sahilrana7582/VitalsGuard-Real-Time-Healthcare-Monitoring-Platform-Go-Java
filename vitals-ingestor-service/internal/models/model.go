package models

import (
	"time"
)

type Vitals struct {
	ID                 string    `json:"id" db:"id"`
	TenantID           string    `json:"tenant_id" db:"tenant_id"`
	PatientID          string    `json:"patient_id" db:"patient_id"`
	HeartRate          int       `json:"heart_rate" db:"heart_rate"`
	BloodPressureUpper int       `json:"blood_pressure_upper" db:"blood_pressure_upper"`
	BloodPressureLower int       `json:"blood_pressure_lower" db:"blood_pressure_lower"`
	RespiratoryRate    int       `json:"respiratory_rate" db:"respiratory_rate"`
	OxygenSaturation   int       `json:"oxygen_saturation" db:"oxygen_saturation"`
	Temperature        float64   `json:"temperature" db:"temperature"`
	GlucoseLevel       *float64  `json:"glucose_level,omitempty" db:"glucose_level"`
	RecordedAt         time.Time `json:"recorded_at" db:"recorded_at"`
	CreatedAt          time.Time `json:"created_at" db:"created_at"`
	UpdatedAt          time.Time `json:"updated_at" db:"updated_at"`
}

type Patient struct {
	ID         string `json:"id" db:"id"`
	TenantID   string `json:"tenant_id" db:"tenant_id"`
	HospitalID int    `json:"hospital_id" db:"hospital_id"`
	Vitals     Vitals `json:"vitals"`
}

type KafkaMessage struct {
	PatientID          string    `json:"patient_id"`
	TenantID           string    `json:"tenant_id"`
	HeartRate          int       `json:"heart_rate"`
	BloodPressureUpper int       `json:"blood_pressure_upper"`
	BloodPressureLower int       `json:"blood_pressure_lower"`
	RespiratoryRate    int       `json:"respiratory_rate"`
	OxygenSaturation   int       `json:"oxygen_saturation"`
	Temperature        float64   `json:"temperature"`
	GlucoseLevel       *float64  `json:"glucose_level,omitempty"`
	RecordedAt         time.Time `json:"recorded_at"`
}

func NewKafkaMessage(p Patient) KafkaMessage {
	return KafkaMessage{
		PatientID:          p.ID,
		TenantID:           p.TenantID,
		HeartRate:          p.Vitals.HeartRate,
		BloodPressureUpper: p.Vitals.BloodPressureUpper,
		BloodPressureLower: p.Vitals.BloodPressureLower,
		RespiratoryRate:    p.Vitals.RespiratoryRate,
		OxygenSaturation:   p.Vitals.OxygenSaturation,
		Temperature:        p.Vitals.Temperature,
		GlucoseLevel:       p.Vitals.GlucoseLevel,
		RecordedAt:         p.Vitals.RecordedAt,
	}
}
