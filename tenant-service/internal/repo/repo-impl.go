package repo

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"
	"vitals-guard/common/errs"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v5/pgxpool"
)

type repo struct {
	db *pgxpool.Pool
}

func NewTenantRepo(db *pgxpool.Pool) TenantRepository {
	return &repo{db: db}
}

func (r *repo) CreateTenant(ctx context.Context, dto *TenantCreateDTO) (*Tenant, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	const query = `
		INSERT INTO tenants (name, code, email, phone)
		VALUES ($1, $2, $3, $4)
		RETURNING id, name, code, email, phone, created_at, updated_at;
	`

	var tenant Tenant
	err := r.db.QueryRow(ctx, query,
		dto.Name,
		dto.Code,
		dto.Email,
		dto.Phone,
	).Scan(
		&tenant.ID,
		&tenant.Name,
		&tenant.Code,
		&tenant.Email,
		&tenant.Phone,
		&tenant.CreatedAt,
		&tenant.UpdatedAt,
	)

	if err != nil {
		var pgErr *pgx.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			case "23505":
				return nil, errs.New(fmt.Sprintf("Tenant with name or code already exists: %s", pgErr.Detail), "DUPLICATE_ERROR", http.StatusBadRequest)
			case "23514":
				return nil, errs.New(fmt.Sprintf("Validation failed: %s", pgErr.Detail), "DB_VALIDATION_ERROR", http.StatusBadRequest)
			default:
				return nil, errs.ErrInternalServer
			}
		}
		return nil, errs.ErrInternalServer
	}

	return &tenant, nil
}
