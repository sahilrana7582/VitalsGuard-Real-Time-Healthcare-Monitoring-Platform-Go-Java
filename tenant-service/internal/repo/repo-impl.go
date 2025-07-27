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
				return nil, errs.New("Internal server error", "INTERNAL_ERROR", http.StatusInternalServerError)
			}
		}
		return nil, errs.New("Internal server error", err.Error(), http.StatusInternalServerError)
	}

	return &tenant, nil
}

func (r *repo) CreateTenantProfile(ctx context.Context, profileDTO TenantProfileCreateDTO) (*TenantProfile, error) {
	var tenantProfile TenantProfile

	query := `
		INSERT INTO tenant_profiles (
			tenant_id, legal_name, address,
			city, state, country,
			postal_code, gst_number, license_number, emergency_contact
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		RETURNING 
			id, tenant_id, legal_name, address, 
			city, state, country, postal_code, 
			gst_number, license_number, emergency_contact,
			created_at, updated_at;
	`

	row := r.db.QueryRow(ctx, query,
		profileDTO.TenantID,
		profileDTO.LegalName,
		profileDTO.Address,
		profileDTO.City,
		profileDTO.State,
		profileDTO.Country,
		profileDTO.PostalCode,
		profileDTO.GSTNumber,
		profileDTO.LicenseNumber,
		profileDTO.EmergencyContact,
	)

	err := row.Scan(
		&tenantProfile.ID,
		&tenantProfile.TenantID,
		&tenantProfile.LegalName,
		&tenantProfile.Address,
		&tenantProfile.City,
		&tenantProfile.State,
		&tenantProfile.Country,
		&tenantProfile.PostalCode,
		&tenantProfile.GSTNumber,
		&tenantProfile.LicenseNumber,
		&tenantProfile.EmergencyContact,
		&tenantProfile.CreatedAt,
		&tenantProfile.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, context.Canceled) {
			return nil, fmt.Errorf("request canceled: %w", err)
		}
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, fmt.Errorf("request timed out: %w", err)
		}
		return nil, fmt.Errorf("failed to insert tenant profile: %w", err)
	}

	return &tenantProfile, nil
}
