package repo

type TenantCreateDTO struct {
	Name  string
	Code  string
	Email string
	Phone *string
}

type TenantProfileCreateDTO struct {
	TenantID         string
	LegalName        string
	Address          string
	City             string
	State            string
	Country          string
	PostalCode       string
	GSTNumber        *string
	LicenseNumber    *string
	LogoURL          *string
	Website          *string
	EmergencyContact *string
}
