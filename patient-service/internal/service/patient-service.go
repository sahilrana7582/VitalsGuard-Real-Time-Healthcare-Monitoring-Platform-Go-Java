package service

import (
	"context"

	"github.com/sahilrana7582/vitals-guard/patient-service/internal/dto"
)

type IPatientService interface {
	CreatePatient(ctx context.Context, tenantID, userID string, p dto.NewPatientDTO) (*dto.PatientDTOResponse, error)
	CreatePatientProfile(ctx context.Context, tenantID, userID string, p dto.NewPatientProfileDTO) (*dto.PatientProfileDTOResponse, error)
	CreateVitalsInfo(ctx context.Context, tenantID, userID string, v dto.NewVitalsInfoDTO) (*dto.VitalsInfoDTOResponse, error)
}
