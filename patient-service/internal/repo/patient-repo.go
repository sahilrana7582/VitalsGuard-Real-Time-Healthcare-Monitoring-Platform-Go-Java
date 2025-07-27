package repo

import (
	"context"

	"github.com/sahilrana7582/vitals-guard/patient-service/internal/dto"
)

type IPatientRepo interface {
	CreatePatient(ctx context.Context, payload dto.NewPatientDTO) error

	CreatePatientProfile(ctx context.Context, payload dto.NewPatientProfileDTO) error

	CreateVitalsInfo(ctx context.Context, payload dto.NewVitalsInfoDTO) error
}
