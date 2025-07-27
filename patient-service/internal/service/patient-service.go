package service

import (
	"context"

	"github.com/sahilrana7582/vitals-guard/patient-service/internal/dto"
)

type IPatientService interface {
	CreatePatient(ctx context.Context, p dto.NewPatientDTO) (*dto.PatientDTOResponse, error)
	CreatePatientProfile(ctx context.Context, p dto.NewPatientProfileDTO) (*dto.PatientProfileDTOResponse, error)
	CreateVitalsInfo(ctx context.Context, v dto.NewVitalsInfoDTO) (*dto.VitalsInfoDTOResponse, error)
}
