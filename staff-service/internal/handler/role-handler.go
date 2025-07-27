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

func (h *RoleHandler) AssignUserRole(w http.ResponseWriter, r *http.Request) error {
	tenantID := apicommon.GetTenantID(r)

	if tenantID == "" {
		return errs.New("BAD_REQUEST", "NO TENANT ID", http.StatusBadRequest)
	}

	userID := apicommon.ReadParam(r, "userID")
	if userID == "" {
		return errs.New("BAD_REQUEST", "NO User ID", http.StatusBadRequest)
	}

	roleID := apicommon.ReadParam(r, "roleID")
	if roleID == "" {
		return errs.New("BAD_REQUEST", "NO Role ID", http.StatusBadRequest)
	}

	resp, err := h.service.AssignRole(r.Context(), tenantID, userID, roleID)

	if err != nil {
		return err
	}

	return apicommon.WriteSuccess(w, http.StatusCreated, "Role assigned successfully", resp)

}
