package handler

import (
	"encoding/json"
	"net/http"
	apicommon "vitals-guard/common/api-common"
	"vitals-guard/common/errs"

	"github.com/sahilrana7582/vitals-guard/staff-service/internal/dto"
	service "github.com/sahilrana7582/vitals-guard/staff-service/internal/service/staff"
)

type StaffHandler struct {
	service service.IStaffService
}

func NewStaffHandler(s service.IStaffService) *StaffHandler {
	return &StaffHandler{service: s}
}

func (h *StaffHandler) CreateStaff(w http.ResponseWriter, r *http.Request) error {
	var payload dto.NewStaffDTO

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		return errs.New("BAD_BODY", "Invalid JSON payload", http.StatusBadRequest)
	}
	defer r.Body.Close()

	resp, err := h.service.CreateStaff(r.Context(), payload)

	if err != nil {
		return err
	}

	return apicommon.WriteSuccess(w, http.StatusCreated, "New staff addeed", resp)
}

func (h *StaffHandler) CreateDoctor(w http.ResponseWriter, r *http.Request) error {
	var payload dto.NewDoctorDTO
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		return errs.New("BAD_REQUEST", "Invalid Request Body", http.StatusBadRequest)
	}
	defer r.Body.Close()

	resp, err := h.service.RegisterDoctor(r.Context(), payload)

	if err != nil {
		return err
	}

	return apicommon.WriteSuccess(w, http.StatusCreated, "Doctore register successfully", resp)
}

func (h *StaffHandler) CreateNurse(w http.ResponseWriter, r *http.Request) error {
	var payload dto.NewNurseDTO
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		return errs.New("BAD_REQUEST", "Invalid Request Body", http.StatusBadRequest)

	}
	defer r.Body.Close()

	resp, err := h.service.RegisterNurse(r.Context(), payload)

	if err != nil {
		return err
	}

	return apicommon.WriteSuccess(w, http.StatusCreated, resp.Message, resp)
}
