package repo

import (
	"context"
	"net/http"
	"vitals-guard/common/errs"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sahilrana7582/vitals-guard/staff-service/internal/dto"
)

type staffRepo struct {
	db *pgxpool.Pool
}

func NewStaffRepo(db *pgxpool.Pool) IStaffRepo {
	return &staffRepo{
		db: db,
	}
}

func (r *staffRepo) CreateStaff(ctx context.Context, payload dto.NewStaffDTO) error {
	query := `
		INSERT INTO staff (
			tenant_id, user_id, full_name, gender, dob, contact_number, email, address
		)
		VALUES (
			 $1, $2, $3, $4, $5, $6, $7, $8
		)
	`

	_, err := r.db.Exec(ctx, query,
		payload.TenantID,
		payload.UserID,
		payload.FullName,
		payload.Gender,
		payload.DOB,
		payload.ContactNumber,
		payload.Email,
		payload.Address,
	)

	if err != nil {
		return errs.New("DATABASE_ERROR", "Error in SQL: "+err.Error(), http.StatusInternalServerError)
	}

	return nil
}
