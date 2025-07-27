package service

import (
	"context"

	"github.com/sahilrana7582/vitals-guard/staff-service/internal/dto"
)

type IRoleService interface {
	CreateRole(ctx context.Context, tenantID string, payload *dto.NewRoleDTO) (*dto.RoleDTOResponse, error)
	// GetRoleByID(ctx context.Context, id string) (*models.Role, error)
	// GetAllRoles(ctx context.Context, tenantID string) ([]*models.Role, error)
	// DeleteRole(ctx context.Context, id string) error
	// UpdateRole(ctx context.Context, id string, payload *dto.UpdateRoleDTO) (*models.Role, error)
}
