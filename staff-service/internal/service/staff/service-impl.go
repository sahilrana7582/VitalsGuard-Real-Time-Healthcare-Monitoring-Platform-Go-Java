package service

import (
	"context"
	"fmt"

	"github.com/sahilrana7582/vitals-guard/staff-service/internal/dto"
	repo "github.com/sahilrana7582/vitals-guard/staff-service/internal/repo/staff"
)

type staffService struct {
	repo repo.IStaffRepo
}

func NewStaffService(repo repo.IStaffRepo) IStaffService {
	return &staffService{repo: repo}
}

func (s *staffService) CreateStaff(ctx context.Context, input dto.NewStaffDTO) (*dto.StaffDTOResponse, error) {

	err := s.repo.CreateStaff(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("failed to create staff: %w", err)
	}

	return &dto.StaffDTOResponse{Message: "Staff created successfully"}, nil
}

// func (s *staffService) RegisterDoctor(ctx context.Context, input dto.NewDoctorDTO) (*dto.DoctorDTOResponse, error) {
// 	if input.TenantID == "" || input.StaffID == "" || input.Specialization == "" {
// 		return nil, errors.New("tenant_id, staff_id, and specialization are required fields")
// 	}

// 	err := s.repo.CreateDoctor(ctx, input)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to register doctor: %w", err)
// 	}

// 	return &dto.DoctorDTOResponse{Message: "Doctor registered successfully"}, nil
// }

// func (s *staffService) RegisterNurse(ctx context.Context, input dto.NewNurseDTO) (*dto.NurseDTOResponse, error) {
// 	if input.TenantID == "" || input.StaffID == "" || input.Shift == "" {
// 		return nil, errors.New("tenant_id, staff_id, and shift are required fields")
// 	}

// 	err := s.repo.CreateNurse(ctx, input)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to register nurse: %w", err)
// 	}

// 	return &dto.NurseDTOResponse{Message: "Nurse registered successfully"}, nil
// }
