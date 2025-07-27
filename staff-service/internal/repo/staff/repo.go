package repo

import (
	"context"

	"github.com/sahilrana7582/vitals-guard/staff-service/internal/dto"
)

type IStaffRepo interface {
	CreateStaff(ctx context.Context, payload dto.NewStaffDTO) error
	// CreateDoctor(ctx context.Context, payload dto.NewDoctorDTO) error
	// CreateNurse(ctx context.Context, payload dto.NewNurseDTO) error
}
