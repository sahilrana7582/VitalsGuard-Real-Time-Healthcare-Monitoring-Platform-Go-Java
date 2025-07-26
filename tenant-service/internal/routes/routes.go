package routes

import (
	"encoding/json"
	"net/http"
	"time"

	apicommon "vitals-guard/common/api-common"

	"github.com/go-chi/chi/v5"
	"github.com/sahilrana7582/vitals-guard/tenant-service/internal/handler"
)

func TenantRoutes(h *handler.TenantHandler) http.Handler {
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
		r.Post("/", apicommon.ErrorHandler(h.CreateTenant))
	})

	return r
}
