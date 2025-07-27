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

func (r *patientService) CreatePatient(ctx context.Context, tenantID, userID string, p dto.NewPatientDTO) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return errs.New("DB_TX_BEGIN", fmt.Sprintf("repo: begin transaction failed: %s", err), http.StatusInternalServerError)
	}
	defer tx.Rollback(ctx)

	// Properly set session-local variables via set_config
	if _, err = tx.Exec(ctx, `SELECT set_config('app.user_id', $1, true);`, userID); err != nil {
		return errs.New("DB_CTX_SET", fmt.Sprintf("repo: set user_id failed: %s", err), http.StatusInternalServerError)
	}
	if _, err = tx.Exec(ctx, `SELECT set_config('app.tenant_id', $1, true);`, tenantID); err != nil {
		return errs.New("DB_CTX_SET", fmt.Sprintf("repo: set tenant_id failed: %s", err), http.StatusInternalServerError)
	}

	const q = `
		INSERT INTO patients (tenant_id, full_name, age, admission_reason)
		VALUES ($1, $2, $3, $4)
	`
	if _, err = tx.Exec(ctx, q,
		tenantID,
		p.FullName,
		p.Age,
		p.AdmissionReason,
	); err != nil {
		return errs.New("DB_INSERT", fmt.Sprintf("repo: CreatePatient failed: %s", err), http.StatusInternalServerError)
	}

	if err = tx.Commit(ctx); err != nil {
		return errs.New("DB_TX_COMMIT", fmt.Sprintf("repo: commit transaction failed: %s", err), http.StatusInternalServerError)
	}
	return nil
}

func (r *patientService) CreatePatientProfile(ctx context.Context, tenantID, userID string, p dto.NewPatientProfileDTO) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return errs.New("DB_TX_BEGIN", fmt.Sprintf("repo: begin transaction failed: %s", err), http.StatusInternalServerError)
	}
	defer tx.Rollback(ctx)

	if _, err = tx.Exec(ctx, `SELECT set_config('app.user_id', $1, true);`, userID); err != nil {
		return errs.New("DB_CTX_SET", fmt.Sprintf("repo: set user_id failed: %s", err), http.StatusInternalServerError)
	}
	if _, err = tx.Exec(ctx, `SELECT set_config('app.tenant_id', $1, true);`, tenantID); err != nil {
		return errs.New("DB_CTX_SET", fmt.Sprintf("repo: set tenant_id failed: %s", err), http.StatusInternalServerError)
	}

	dob, err := time.Parse("2006-01-02", p.DOB)
	if err != nil {
		return errs.New("INVALID_INPUT", fmt.Sprintf("repo: invalid DOB format: %s", err), http.StatusBadRequest)
	}

	const q = `
		INSERT INTO patient_tables
			(tenant_id, patient_id, gender, dob, blood_group, contact_number, email, address, postal_code)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`
	if _, err = tx.Exec(ctx, q,
		tenantID,
		p.PatientID,
		p.Gender,
		dob,
		p.BloodGroup,
		p.ContactNumber,
		p.Email,
		p.Address,
		p.PostalCode,
	); err != nil {
		return errs.New("DB_INSERT", fmt.Sprintf("repo: CreatePatientProfile failed: %s", err), http.StatusInternalServerError)
	}

	if err = tx.Commit(ctx); err != nil {
		return errs.New("DB_TX_COMMIT", fmt.Sprintf("repo: commit transaction failed: %s", err), http.StatusInternalServerError)
	}
	return nil
}

func (r *patientService) CreateVitalsInfo(ctx context.Context, tenantID, userID string, v dto.NewVitalsInfoDTO) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return errs.New("DB_TX_BEGIN", fmt.Sprintf("repo: begin transaction failed: %s", err), http.StatusInternalServerError)
	}
	defer tx.Rollback(ctx)

	if _, err = tx.Exec(ctx, `SELECT set_config('app.user_id', $1, true);`, userID); err != nil {
		return errs.New("DB_CTX_SET", fmt.Sprintf("repo: set user_id failed: %s", err), http.StatusInternalServerError)
	}
	if _, err = tx.Exec(ctx, `SELECT set_config('app.tenant_id', $1, true);`, tenantID); err != nil {
		return errs.New("DB_CTX_SET", fmt.Sprintf("repo: set tenant_id failed: %s", err), http.StatusInternalServerError)
	}

	recordedAt, err := time.Parse(time.RFC3339, v.RecordedAt)
	if err != nil {
		return errs.New("INVALID_INPUT", fmt.Sprintf("repo: invalid recorded_at format: %s", err), http.StatusBadRequest)
	}

	const q = `
		INSERT INTO vitals_info
			(tenant_id, patient_id,
			 heart_rate, blood_pressure_upper, blood_pressure_lower,
			 respiratory_rate, oxygen_saturation, temperature, glucose_level, recorded_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`
	if _, err = tx.Exec(ctx, q,
		tenantID,
		v.PatientID,
		v.HeartRate,
		v.BloodPressureUpper,
		v.BloodPressureLower,
		v.RespiratoryRate,
		v.OxygenSaturation,
		v.Temperature,
		v.GlucoseLevel,
		recordedAt,
	); err != nil {
		return errs.New("DB_INSERT", fmt.Sprintf("repo: CreateVitalsInfo failed: %s", err), http.StatusInternalServerError)
	}

	if err = tx.Commit(ctx); err != nil {
		return errs.New("DB_TX_COMMIT", fmt.Sprintf("repo: commit transaction failed: %s", err), http.StatusInternalServerError)
	}
	return nil
}
