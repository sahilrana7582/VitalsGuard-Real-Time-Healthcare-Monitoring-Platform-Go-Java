package repo

import (
	"context"
	"fmt"
	"net/http"
	"time"
	"vitals-guard/common/errs"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sahilrana7582/vitals-guard/patient-service/internal/dto"
)

type patientService struct {
	db *pgxpool.Pool
}

func NewPatientRepo(db *pgxpool.Pool) IPatientRepo {
	return &patientService{db: db}
}

func (r *patientService) CreatePatient(ctx context.Context, p dto.NewPatientDTO) error {
	const q = `
		INSERT INTO patients (tenant_id, full_name, age, admission_reason)
		VALUES ($1, $2, $3, $4)
	`
	_, err := r.db.Exec(ctx, q,
		p.TenantID,
		p.FullName,
		p.Age,
		p.AdmissionReason,
	)
	if err != nil {
		return errs.New("DB_ERROR", fmt.Sprintf("repo: CreatePatient failed: %w", err.Error()), http.StatusInternalServerError)
	}
	return nil
}

func (r *patientService) CreatePatientProfile(ctx context.Context, p dto.NewPatientProfileDTO) error {
	const q = `
		INSERT INTO patient_tables
			(tenant_id, patient_id, gender, dob, blood_group, contact_number, email, address, postal_code)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`
	dob, err := time.Parse("2006-01-02", p.DOB)
	if err != nil {
		return errs.New("DB_ERROR", fmt.Sprintf("repo: Create_Patient_Profile failed: %w", err.Error()), http.StatusInternalServerError)

	}

	_, err = r.db.Exec(ctx, q,
		p.TenantID,
		p.PatientID,
		p.Gender,
		dob,
		p.BloodGroup,
		p.ContactNumber,
		p.Email,
		p.Address,
		p.PostalCode,
	)
	if err != nil {
		return fmt.Errorf("repo: CreatePatientProfile failed: %w", err)
	}
	return nil
}

func (r *patientService) CreateVitalsInfo(ctx context.Context, v dto.NewVitalsInfoDTO) error {
	const q = `
		INSERT INTO vitals_info
			(tenant_id, patient_id,
			 heart_rate, blood_pressure_upper, blood_pressure_lower,
			 respiratory_rate, oxygen_saturation, temperature, glucose_level, recorded_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`

	recordedAt, err := time.Parse(time.RFC3339, v.RecordedAt)
	if err != nil {
		return errs.New("DB_ERROR", fmt.Sprintf("repo: Create_Vital failed: %w", err.Error()), http.StatusInternalServerError)

	}

	_, err = r.db.Exec(ctx, q,
		v.TenantID,
		v.PatientID,
		v.HeartRate,
		v.BloodPressureUpper,
		v.BloodPressureLower,
		v.RespiratoryRate,
		v.OxygenSaturation,
		v.Temperature,
		v.GlucoseLevel,
		recordedAt,
	)
	if err != nil {
		return fmt.Errorf("repo: CreateVitalsInfo failed: %w", err)
	}
	return nil
}
