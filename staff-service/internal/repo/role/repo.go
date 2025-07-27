package repo

import (
	"context"

	"github.com/sahilrana7582/vitals-guard/staff-service/internal/dto"
)

type IRoleRepo interface {
	CreateRole(ctx context.Context, role *dto.NewRoleDTO) error
	AssignRole(ctx context.Context, tenantID, userID, roleID string) error
	// GetRoleByID(ctx context.Context, id string) (*models.Role, error)
	// GetRolesByTenantID(ctx context.Context, tenantID string) ([]*models.Role, error)
	// UpdateRole(ctx context.Context, role *models.Role) error
	// DeleteRole(ctx context.Context, id string) error
}
