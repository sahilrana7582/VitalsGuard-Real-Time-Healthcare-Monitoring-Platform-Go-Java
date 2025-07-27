package dto

type NewStaffDTO struct {
	TenantID      string `json:"tenant_id"`
	UserID        string `json:"user_id"`
	FullName      string `json:"full_name"`
	Gender        string `json:"gender"`
	DOB           string `json:"dob"`
	ContactNumber string `json:"contact_number"`
	Email         string `json:"email"`
	Address       string `json:"address"`
	Department    string `json:"department"`
}

type StaffDTOResponse struct {
	Message string `json:"message"`
}

type NewDoctorDTO struct {
	TenantID          string `json:"tenant_id"`
	StaffID           string `json:"staff_id"`
	Specialization    string `json:"specialization"`
	LicenseNumber     string `json:"license_number"`
	YearsOfExperience int    `json:"years_of_experience"`
}

type DoctorDTOResponse struct {
	Message string `json:"message"`
}

type NewNurseDTO struct {
	TenantID      string `json:"tenant_id"`
	StaffID       string `json:"staff_id"`
	Shift         string `json:"shift"`
	FloorAssigned string `json:"floor_assigned"`
}

type NurseDTOResponse struct {
	Message string `json:"message"`
}
