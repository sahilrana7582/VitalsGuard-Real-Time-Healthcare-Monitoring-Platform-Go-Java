package handler

import (
	"encoding/json"
	"net/http"
	apicommon "vitals-guard/common/api-common"
	"vitals-guard/common/errs"

	"github.com/go-chi/chi/v5"
	"github.com/sahilrana7582/vitals-guard/patient-service/internal/dto"
	"github.com/sahilrana7582/vitals-guard/patient-service/internal/service"
)

type PatientHandler struct {
	service service.IPatientService
}

func NewPatientHandler(svc service.IPatientService) *PatientHandler {
	return &PatientHandler{service: svc}
}

func (h *PatientHandler) CreatePatient(w http.ResponseWriter, r *http.Request) error {
	var payload dto.NewPatientDTO
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		return errs.New("BAD_REQUEST", "INVALID REQUEST BODY", http.StatusBadRequest)
	}
	defer r.Body.Close()

	tenantID := apicommon.GetTenantID(r)
	if tenantID == "" {
		return errs.New("BAD_REQUEST", "INVALID TENANT ID", http.StatusBadRequest)

	}

	userID := apicommon.GetUserID(r)
	if userID == "" {
		return errs.New("BAD_REQUEST", "INVALID USER ID", http.StatusBadRequest)

	}

	resp, err := h.service.CreatePatient(r.Context(), tenantID, userID, payload)
	if err != nil {
		return err
	}
	return apicommon.WriteSuccess(w, http.StatusCreated, resp.Message, resp)
}

func (h *PatientHandler) CreatePatientProfile(w http.ResponseWriter, r *http.Request) error {
	patientID := chi.URLParam(r, "patientID")
	if patientID == "" {
		return errs.New("BAD_REQUEST", "INVALID REQUEST BODY", http.StatusBadRequest)

	}

	var payload dto.NewPatientProfileDTO
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		return errs.New("BAD_REQUEST", "INVALID REQUEST BODY", http.StatusBadRequest)
	}
	defer r.Body.Close()

	tenantID := apicommon.GetTenantID(r)
	if tenantID == "" {
		return errs.New("BAD_REQUEST", "INVALID TENANT ID", http.StatusBadRequest)

	}

	userID := apicommon.GetUserID(r)
	if userID == "" {
		return errs.New("BAD_REQUEST", "INVALID USER ID", http.StatusBadRequest)

	}

	payload.PatientID = patientID
	resp, err := h.service.CreatePatientProfile(r.Context(), tenantID, userID, payload)
	if err != nil {
		return err
	}
	return apicommon.WriteSuccess(w, http.StatusCreated, resp.Message, resp)

}

func (h *PatientHandler) CreateVitalsInfo(w http.ResponseWriter, r *http.Request) error {
	patientID := chi.URLParam(r, "patientID")
	if patientID == "" {
		return errs.New("BAD_REQUEST", "INVALID REQUEST BODY", http.StatusBadRequest)

	}

	var payload dto.NewVitalsInfoDTO
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		return errs.New("BAD_REQUEST", "INVALID REQUEST BODY", http.StatusBadRequest)
	}
	defer r.Body.Close()

	tenantID := apicommon.GetTenantID(r)
	if tenantID == "" {
		return errs.New("BAD_REQUEST", "INVALID TENANT ID", http.StatusBadRequest)

	}

	userID := apicommon.GetUserID(r)
	if userID == "" {
		return errs.New("BAD_REQUEST", "INVALID USER ID", http.StatusBadRequest)

	}

	payload.PatientID = patientID
	resp, err := h.service.CreateVitalsInfo(r.Context(), tenantID, userID, payload)
	if err != nil {
		return err
	}
	return apicommon.WriteSuccess(w, http.StatusCreated, resp.Message, resp)

}
