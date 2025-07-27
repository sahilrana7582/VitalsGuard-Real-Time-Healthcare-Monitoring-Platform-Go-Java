package repo

import (
	"context"
	"errors"
	"net/http"
	"time"
	"vitals-guard/common/errs"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sahilrana7582/vitals-guard/staff-service/internal/dto"
)

type roleRepo struct {
	db *pgxpool.Pool
}

func NewRoleRepo(db *pgxpool.Pool) IRoleRepo {
	return &roleRepo{
		db: db,
	}
}

func (r *roleRepo) CreateRole(ctx context.Context, role *dto.NewRoleDTO) error {
	query := `
		INSERT INTO roles (tenant_id, name, description, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`

	now := time.Now()

	var insertedID string
	err := r.db.QueryRow(ctx, query,
		role.TenantID,
		role.Name,
		role.Description,
		now,
		now,
	).Scan(&insertedID)

	if err != nil {
		if errors.Is(err, context.Canceled) {
			return errs.New("REQUEST_TIMEOUT", "request cancelled by client", http.StatusRequestTimeout)
		}
		if pgErr, ok := err.(*pgx.PgError); ok {
			switch pgErr.Code {
			case "23505":
				return errs.New("DUPLICATE_KEY", "role already exists", http.StatusBadRequest)
			case "23503":
				return errs.New("INVALID_TENANT_KEY", "invalid foreign key: tenant does not exist", http.StatusBadRequest)
			default:
				return errs.New("Query_Failed", "failed to insert role: "+pgErr.Message, http.StatusInternalServerError)
			}
		}
		return errs.ErrInternalServer
	}

	return nil
}

func (r *roleRepo) AssignRole(ctx context.Context, tenantID, userID, roleID string) error {
	query := `
		INSERT INTO user_roles (user_id, role_id)
		SELECT u.id, $2
		FROM users u
		WHERE u.id = $1 AND u.tenant_id = $3
		ON CONFLICT DO NOTHING
	`

	tag, err := r.db.Exec(ctx, query, userID, roleID, tenantID)
	if err != nil {
		return err
	}

	if tag.RowsAffected() == 0 {
		return errs.New("NO_INSERTION", "Maybe either user or role does not exits", http.StatusBadRequest)
	}

	return nil
}
