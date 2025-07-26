package service

import (
	"context"

	"github.com/sahilrana7582/vitals-guard/tenant-service/internal/repo"
)

type ITenantService interface {
	CreateTenant(ctx context.Context, dto *repo.TenantCreateDTO) (*repo.Tenant, error)
}
