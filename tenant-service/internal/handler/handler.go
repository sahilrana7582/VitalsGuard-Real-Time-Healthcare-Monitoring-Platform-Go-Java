package handler

import (
	"encoding/json"
	"net/http"
	apicommon "vitals-guard/common/api-common"
	"vitals-guard/common/errs"

	"github.com/sahilrana7582/vitals-guard/tenant-service/internal/repo"
	"github.com/sahilrana7582/vitals-guard/tenant-service/internal/service"
)

type TenantHandler struct {
	service service.ITenantService
}

func NewTenantHandler(s service.ITenantService) *TenantHandler {
	return &TenantHandler{service: s}
}

func (h *TenantHandler) CreateTenant(w http.ResponseWriter, r *http.Request) error {
	var dto repo.TenantCreateDTO

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		return errs.New("Failed to parse request body", "INVALID_JSON", http.StatusBadRequest)
	}

	tenant, appErr := h.service.CreateTenant(r.Context(), &dto)
	if appErr != nil {
		return appErr
	}

	return apicommon.WriteSuccess(w, http.StatusCreated, "Tenant created successfully", tenant)
}
