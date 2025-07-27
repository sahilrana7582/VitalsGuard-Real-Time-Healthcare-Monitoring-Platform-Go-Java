package handler

import (
	"encoding/json"
	"net/http"
	apicommon "vitals-guard/common/api-common"
	"vitals-guard/common/errs"

	"github.com/sahilrana7582/vitals-guard/staff-service/internal/dto"
	service "github.com/sahilrana7582/vitals-guard/staff-service/internal/service/role"
)

type RoleHandler struct {
	service service.IRoleService
}

func NewRoleHandler(svc service.IRoleService) *RoleHandler {
	return &RoleHandler{
		service: svc,
	}
}

func (h *RoleHandler) CreateRole(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	var payload dto.NewRoleDTO
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		return err
	}

	tenantID := apicommon.GetTenantID(r)

	if tenantID == "" {
		return errs.New("BAD_REQUEST", "Tenant id not provided", http.StatusBadRequest)
	}
	response, err := h.service.CreateRole(ctx, tenantID, &payload)
	if err != nil {
		return err
	}

	return apicommon.WriteSuccess(w, http.StatusCreated, "Role created successfully", response)
}
