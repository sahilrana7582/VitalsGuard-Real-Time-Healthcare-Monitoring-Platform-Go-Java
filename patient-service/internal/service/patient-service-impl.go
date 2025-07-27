package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/sahilrana7582/vitals-guard/patient-service/internal/dto"
	"github.com/sahilrana7582/vitals-guard/patient-service/internal/repo"
)

type patientService struct {
	repo repo.IPatientRepo
}

func NewPatientService(r repo.IPatientRepo) IPatientService {
	return &patientService{repo: r}
}

func (s *patientService) CreatePatient(ctx context.Context, tenantID, userID string, p dto.NewPatientDTO) (*dto.PatientDTOResponse, error) {

	if err := s.repo.CreatePatient(ctx, tenantID, userID, p); err != nil {
		return nil, fmt.Errorf("failed to create patient: %w", err)
	}

	return &dto.PatientDTOResponse{Message: "Patient created successfully"}, nil
}

func (s *patientService) CreatePatientProfile(ctx context.Context, tenantID, userID string, p dto.NewPatientProfileDTO) (*dto.PatientProfileDTOResponse, error) {

	if _, err := time.Parse("2006-01-02", p.DOB); err != nil {
		return nil, errors.New("dob must be in YYYY-MM-DD format")
	}
	if err := s.repo.CreatePatientProfile(ctx, tenantID, userID, p); err != nil {
		return nil, fmt.Errorf("failed to create patient profile: %w", err)
	}
	return &dto.PatientProfileDTOResponse{Message: "Patient profile created successfully"}, nil
}

func (s *patientService) CreateVitalsInfo(ctx context.Context, tenantID, userID string, v dto.NewVitalsInfoDTO) (*dto.VitalsInfoDTOResponse, error) {

	if _, err := time.Parse(time.RFC3339, v.RecordedAt); err != nil {
		return nil, errors.New("recorded_at must be in RFC3339 format")
	}
	if err := s.repo.CreateVitalsInfo(ctx, tenantID, userID, v); err != nil {
		return nil, fmt.Errorf("failed to create vitals info: %w", err)
	}
	return &dto.VitalsInfoDTOResponse{Message: "Vitals info recorded successfully"}, nil
}
