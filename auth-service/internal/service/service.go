package service

import (
	"context"

	"github.com/sahilrana7582/vitals-guard/auth-service/internal/dto"
)

type IAuthService interface {
	SignUp(ctx context.Context, req *dto.SignUpRequest) (*dto.SignUpResponse, error)
	// Login(ctx context.Context, req *dto.LoginRequest) (*dto.LoginResponse, error)
}
