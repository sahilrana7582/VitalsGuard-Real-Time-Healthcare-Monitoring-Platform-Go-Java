package apicommon

import "net/http"

func GetTenantID(r *http.Request) string {
	tenantID := r.Header.Get("X-Tenant-ID")
	return tenantID
}

func GetUserID(r *http.Request) string {
	tenantID := r.Header.Get("X-User-ID")
	return tenantID
}
