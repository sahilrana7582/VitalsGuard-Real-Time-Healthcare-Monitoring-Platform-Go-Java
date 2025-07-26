package handler

import (
	"encoding/json"
	"net/http"
	apicommon "vitals-guard/common/api-common"
	"vitals-guard/common/errs"

	"github.com/sahilrana7582/vitals-guard/auth-service/internal/dto"
	"github.com/sahilrana7582/vitals-guard/auth-service/internal/service"
)

type AuthHandler struct {
	service service.IAuthService
}

func NewAuthHandler(service service.IAuthService) *AuthHandler {
	return &AuthHandler{
		service: service,
	}
}

func (h *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) error {
	var req dto.SignUpRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return errs.ErrBadRequest
	}

	res, err := h.service.SignUp(r.Context(), &req)
	if err != nil {
		return err
	}

	return apicommon.WriteSuccess(w, http.StatusCreated, "User Signed Up", res)
}
