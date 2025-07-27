package routes

import (
	"encoding/json"
	"net/http"
	"time"
	apicommon "vitals-guard/common/api-common"

	"github.com/go-chi/chi/v5"
	"github.com/sahilrana7582/vitals-guard/staff-service/internal/handler"
)

func NewStaffRoutes(staffHandler *handler.StaffHandler, roleHandler *handler.RoleHandler) http.Handler {
	r := chi.NewRouter()

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":    http.StatusNotFound,
			"error":     "Not Found",
			"message":   "The requested resource could not be found",
			"path":      r.URL.Path,
			"method":    r.Method,
			"timestamp": time.Now().Format(time.RFC3339),
		})
	})

	r.Route("/", func(r chi.Router) {

		r.Route("/members", func(r chi.Router) {
			r.Post("/", apicommon.ErrorHandler(staffHandler.CreateStaff))
			// r.Get("/", apicommon.ErrorHandler(staffHandler.GetAllStaff))
			// r.Get("/{id}", apicommon.ErrorHandler(staffHandler.GetStaffByID))
			// r.Put("/{id}", apicommon.ErrorHandler(staffHandler.UpdateStaff))
			// r.Delete("/{id}", apicommon.ErrorHandler(staffHandler.DeleteStaff))
		})

		r.Route("/roles", func(r chi.Router) {
			r.Post("/", apicommon.ErrorHandler(roleHandler.CreateRole))
			r.Post("/assign-role/{userID}/{roleID}", apicommon.ErrorHandler(roleHandler.AssignUserRole))
			// r.Get("/", apicommon.ErrorHandler(roleHandler.GetAllRoles))
			// r.Get("/{id}", apicommon.ErrorHandler(roleHandler.GetRoleByID))
			// r.Put("/{id}", apicommon.ErrorHandler(roleHandler.UpdateRole))
			// r.Delete("/{id}", apicommon.ErrorHandler(roleHandler.DeleteRole))
		})
	})

	return r
}
