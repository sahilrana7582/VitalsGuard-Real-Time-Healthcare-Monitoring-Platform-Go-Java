package repo

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sahilrana7582/vitals-guard/vitals-ingestor-service/internal/models"
)

type Repo struct {
	db *pgxpool.Pool
}

func NewRepo(db *pgxpool.Pool) *Repo {
	return &Repo{
		db: db,
	}
}

func (r *Repo) FetchVitals(ctx context.Context, offset, limit int) (*[]models.Vitals, error) {
	query := `
		SELECT 
			id,
			tenant_id,
			patient_id,
			heart_rate,
			blood_pressure_upper,
			blood_pressure_lower,
			respiratory_rate,
			oxygen_saturation,
			temperature,
			glucose_level,
			recorded_at,
			created_at,
			updated_at
		FROM vitals_info
		ORDER BY recorded_at DESC
		OFFSET $1 LIMIT $2
	`

	rows, err := r.db.Query(ctx, query, offset, limit)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %w", err)
	}
	defer rows.Close()

	var vitals []models.Vitals

	for rows.Next() {
		var v models.Vitals
		err := rows.Scan(
			&v.ID,
			&v.TenantID,
			&v.PatientID,
			&v.HeartRate,
			&v.BloodPressureUpper,
			&v.BloodPressureLower,
			&v.RespiratoryRate,
			&v.OxygenSaturation,
			&v.Temperature,
			&v.GlucoseLevel,
			&v.RecordedAt,
			&v.CreatedAt,
			&v.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning vitals row: %w", err)
		}
		vitals = append(vitals, v)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("row iteration error: %w", rows.Err())
	}

	return &vitals, nil
}
