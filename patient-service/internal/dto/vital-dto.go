package dto

type NewVitalsInfoDTO struct {
	TenantID           string   `json:"tenant_id"`
	PatientID          string   `json:"patient_id"`
	HeartRate          int      `json:"heart_rate"`
	BloodPressureUpper *int     `json:"blood_pressure_upper,omitempty"`
	BloodPressureLower *int     `json:"blood_pressure_lower,omitempty"`
	RespiratoryRate    *int     `json:"respiratory_rate,omitempty"`
	OxygenSaturation   *int     `json:"oxygen_saturation,omitempty"`
	Temperature        *float32 `json:"temperature,omitempty"`
	GlucoseLevel       *float32 `json:"glucose_level,omitempty"`
	RecordedAt         string   `json:"recorded_at"`
}

type VitalsInfoDTOResponse struct {
	Message string `json:"message"`
}
