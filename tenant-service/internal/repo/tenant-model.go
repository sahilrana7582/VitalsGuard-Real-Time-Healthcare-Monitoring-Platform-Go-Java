package repo

import "time"

type Tenant struct {
	ID        string
	Name      string
	Code      string
	Email     string
	Phone     *string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type TenantProfile struct {
	ID               string
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
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
