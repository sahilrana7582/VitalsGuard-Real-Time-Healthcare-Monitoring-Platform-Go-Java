package repo

import (
	"context"

	"github.com/sahilrana7582/vitals-guard/auth-service/internal/dto"
)

type IAuthRepo interface {
	// Registers a new user (with tenant-aware uniqueness).
	CreateUser(ctx context.Context, req *dto.SignUpRequest) (*dto.SignUpResponse, error)

	// Fetches user for login based on email and tenant scope.
	// GetUserByEmail(ctx context.Context, tenantID, email string) (*models.User, error)

	// Updates user password after validation or reset.
	// UpdatePassword(ctx context.Context, userID, newHashedPassword string) error
}
