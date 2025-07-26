package repo

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"vitals-guard/common/errs"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sahilrana7582/vitals-guard/auth-service/internal/dto"
)

type authRepo struct {
	db *pgxpool.Pool
}

func NewAuthRepo(db *pgxpool.Pool) IAuthRepo {
	return &authRepo{
		db: db,
	}
}

func (r *authRepo) CreateUser(ctx context.Context, req *dto.SignUpRequest) (*dto.SignUpResponse, error) {
	query := `
		INSERT INTO users (tenant_id, email, password_hash, name)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`

	var userID string
	err := r.db.QueryRow(ctx, query,
		req.TenantID,
		req.Email,
		req.Password,
		req.Name,
	).Scan(&userID)

	if err != nil {
		var pgErr *pgx.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			case "23505":
				return nil, errs.New("duplicate email existis in same system", "DUPLICATE_EMAIL", http.StatusBadRequest)
			case "23503":
				return nil, errs.ErrBadRequest
			default:
				return nil, errs.ErrBadRequest
			}
		}
		return nil, errs.ErrInternalServer
	}

	return &dto.SignUpResponse{
		Message: fmt.Sprintf("User created successfully with ID: %s", userID),
	}, nil
}
