package service

import (
	"context"
	"vitals-guard/common/errs"

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
