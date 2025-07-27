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

func (s *staffService) RegisterDoctor(ctx context.Context, input dto.NewDoctorDTO) (*dto.DoctorDTOResponse, error) {

	err := s.repo.CreateDoctor(ctx, input)
	if err != nil {
		return nil, err
	}

	return &dto.DoctorDTOResponse{Message: "Doctor registered successfully"}, nil
}

func (s *staffService) RegisterNurse(ctx context.Context, input dto.NewNurseDTO) (*dto.NurseDTOResponse, error) {

	err := s.repo.CreateNurse(ctx, input)
	if err != nil {
		return nil, err
	}

	return &dto.NurseDTOResponse{Message: "Nurse registered successfully"}, nil
}
