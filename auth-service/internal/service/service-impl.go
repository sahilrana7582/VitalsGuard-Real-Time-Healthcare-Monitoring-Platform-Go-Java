package service

import (
	"context"
	"vitals-guard/common/errs"
	"vitals-guard/common/token"

	"github.com/sahilrana7582/vitals-guard/auth-service/internal/dto"
	"github.com/sahilrana7582/vitals-guard/auth-service/internal/hash"
	"github.com/sahilrana7582/vitals-guard/auth-service/internal/repo"
)

type authService struct {
	repo repo.IAuthRepo
}

func NewAuthService(repo repo.IAuthRepo) IAuthService {
	return &authService{
		repo: repo,
	}
}

func (s *authService) SignUp(ctx context.Context, req *dto.SignUpRequest) (*dto.SignUpResponse, error) {
	if req.Password == "" {
		return nil, errs.ErrBadRequest
	}

	hashedPwd, err := hash.HashPassword(req.Password)
	if err != nil {
		return nil, errs.ErrInternalServer
	}

	req.Password = hashedPwd

	return s.repo.CreateUser(ctx, req)
}

func (s *authService) Login(ctx context.Context, req *dto.LoginRequest) (*dto.LoginResponse, error) {
	user, err := s.repo.Login(ctx, *req)
	if err != nil {
		return nil, errs.ErrUnauthorized
	}

	err = hash.ComparePasswords(user.PasswordHash, req.Password)
	if err != nil {
		return nil, errs.ErrUnauthorized
	}

	token, err := token.GenerateToken(user.TenantID)
	if err != nil {
		return nil, errs.ErrInternalServer
	}

	return &dto.LoginResponse{
		Token:   token,
		Message: "Login successful",
	}, nil
}
