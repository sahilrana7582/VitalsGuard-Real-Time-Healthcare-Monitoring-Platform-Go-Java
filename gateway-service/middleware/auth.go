package middleware

import (
	"fmt"
	"net/http"
	"strings"
	apicommon "vitals-guard/common/api-common"
	"vitals-guard/common/errs"
	"vitals-guard/common/token"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			apicommon.WriteError(w, errs.New("ERR_UNAUTHORIZED", "Authorization header missing", http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		if tokenStr != "" {
			apicommon.WriteError(w, errs.New("ERR_UNAUTHORIZED", "Authorization header missing", http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		jwtClaims, err := token.ParseJWT(tokenStr)
		if err != nil {
			apicommon.WriteError(w, errs.New("ERR_INVALID_TOKEN", fmt.Sprintf("Reason: %s", err.Error()), http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		tenantID := jwtClaims.TenantID

		r.Header.Set("X-Tenant-ID", tenantID)

		next.ServeHTTP(w, r)

	})
}
