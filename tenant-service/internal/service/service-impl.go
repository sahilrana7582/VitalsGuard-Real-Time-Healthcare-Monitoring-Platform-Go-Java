package service

import (
	"context"

	"github.com/sahilrana7582/vitals-guard/tenant-service/internal/repo"
)

type tenantService struct {
	repo repo.TenantRepository
}

func NewTenantService(r repo.TenantRepository) ITenantService {
	return &tenantService{repo: r}
}

func (s *tenantService) CreateTenant(ctx context.Context, dto *repo.TenantCreateDTO) (*repo.Tenant, error) {

	tenant, err := s.repo.CreateTenant(ctx, dto)
	if err != nil {
		return nil, err
	}

	return tenant, nil
}
