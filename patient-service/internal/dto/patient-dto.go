package dto

type NewPatientDTO struct {
	TenantID        string `json:"tenant_id"`
	FullName        string `json:"full_name"`
	Age             int    `json:"age"`
	AdmissionReason string `json:"admission_reason"`
}

type PatientDTOResponse struct {
	Message string `json:"message"`
}

type NewPatientProfileDTO struct {
	TenantID      string `json:"tenant_id"`
	PatientID     string `json:"patient_id"`
	Gender        string `json:"gender"`
	DOB           string `json:"dob"`
	BloodGroup    string `json:"blood_group"`
	ContactNumber string `json:"contact_number"`
	Email         string `json:"email"`
	Address       string `json:"address"`
	PostalCode    string `json:"postal_code"`
}

type PatientProfileDTOResponse struct {
	Message string `json:"message"`
}
