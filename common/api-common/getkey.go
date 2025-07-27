package apicommon

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func GetTenantID(r *http.Request) string {
	tenantID := r.Header.Get("X-Tenant-ID")
	return tenantID
}

func GetUserID(r *http.Request) string {
	tenantID := r.Header.Get("X-User-ID")
	return tenantID
}

func ReadParam(r *http.Request, key string) string {
	return chi.URLParam(r, key)
}

func ReadQuery(r *http.Request, key string) string {
	return r.URL.Query().Get(key)
}
