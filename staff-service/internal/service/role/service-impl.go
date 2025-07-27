package service

import (
	"context"

	"github.com/sahilrana7582/vitals-guard/staff-service/internal/dto"
	repo "github.com/sahilrana7582/vitals-guard/staff-service/internal/repo/role"
)

type roleService struct {
	repo repo.IRoleRepo
}

func NewRoleService(repo repo.IRoleRepo) IRoleService {
	return &roleService{
		repo: repo,
	}
}

func (s *roleService) CreateRole(ctx context.Context, tenantID string, payload *dto.NewRoleDTO) (*dto.RoleDTOResponse, error) {

	role := &dto.NewRoleDTO{
		TenantID:    tenantID,
		Name:        payload.Name,
		Description: payload.Description,
	}

	if err := s.repo.CreateRole(ctx, role); err != nil {
		return nil, err
	}

	return &dto.RoleDTOResponse{
		Message: "Role created successfully",
	}, nil
}

func (s *roleService) AssignRole(ctx context.Context, tenantID, userID, roleID string) (*dto.AssignRoleResponse, error) {

	err := s.repo.AssignRole(ctx, tenantID, userID, roleID)
	if err != nil {
		return nil, err
	}
	return &dto.AssignRoleResponse{
		Message: "Role assigned successfully",
	}, nil
}
