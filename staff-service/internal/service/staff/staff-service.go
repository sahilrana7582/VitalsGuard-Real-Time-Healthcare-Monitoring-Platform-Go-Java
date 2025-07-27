package service

import (
	"context"

	"github.com/sahilrana7582/vitals-guard/staff-service/internal/dto"
)

type IStaffService interface {
	CreateStaff(ctx context.Context, staff dto.NewStaffDTO) (*dto.StaffDTOResponse, error)

	RegisterDoctor(ctx context.Context, doctor dto.NewDoctorDTO) (*dto.DoctorDTOResponse, error)

	RegisterNurse(ctx context.Context, nurse dto.NewNurseDTO) (*dto.NurseDTOResponse, error)
}
