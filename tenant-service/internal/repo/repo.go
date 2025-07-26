package repo

import (
	"context"
)

type TenantRepository interface {
	CreateTenant(ctx context.Context, dto *TenantCreateDTO) (*Tenant, error)
	// GetTenantByID(ctx context.Context, id string) (*Tenant, error)

	CreateTenantProfile(ctx context.Context, dto TenantProfileCreateDTO) (*TenantProfile, error)
	// GetProfileByTenantID(ctx context.Context, tenantID string) (*TenantProfile, error)
}
