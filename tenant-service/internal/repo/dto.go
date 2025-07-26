package repo

type TenantCreateDTO struct {
	Name  string  `json:"name"`
	Code  string  `json:"code"`
	Email string  `json:"email"`
	Phone *string `json:"phone,omitempty"`
}

type TenantProfileCreateDTO struct {
	TenantID         string  `json:"tenant_id"`
	LegalName        string  `json:"legal_name"`
	Address          string  `json:"address"`
	City             string  `json:"city"`
	State            string  `json:"state"`
	Country          string  `json:"country"`
	LicenseNumber    string  `json:"license_number"`
	PostalCode       string  `json:"postal_code"`
	GSTNumber        *string `json:"gst_number,omitempty"`
	PANNumber        *string `json:"pan_number,omitempty"`
	EmergencyContact *string `json:"emergency_contact,omitempty"`
}
