package repo

import (
	"context"

	"github.com/sahilrana7582/vitals-guard/patient-service/internal/dto"
)

type IPatientRepo interface {
	CreatePatient(ctx context.Context, tenantID, userID string, payload dto.NewPatientDTO) error

	CreatePatientProfile(ctx context.Context, tenantID, userID string, payload dto.NewPatientProfileDTO) error

	CreateVitalsInfo(ctx context.Context, tenantID, userID string, payload dto.NewVitalsInfoDTO) error
}
