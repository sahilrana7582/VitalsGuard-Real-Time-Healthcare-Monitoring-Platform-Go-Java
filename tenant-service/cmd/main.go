package main

import (
	"log"
	"net/http"
	"os"
	"time"

	apicommon "vitals-guard/common/api-common"

	"github.com/sahilrana7582/vitals-guard/tenant-service/config"
	"github.com/sahilrana7582/vitals-guard/tenant-service/internal/db"
	"github.com/sahilrana7582/vitals-guard/tenant-service/internal/handler"
	"github.com/sahilrana7582/vitals-guard/tenant-service/internal/repo"
	"github.com/sahilrana7582/vitals-guard/tenant-service/internal/routes"
	"github.com/sahilrana7582/vitals-guard/tenant-service/internal/service"
)

func main() {
	cfg := config.LoadTenantsConfig(".env")

	dbPool := db.NewConnPool(cfg)

	tenantRepo := repo.NewTenantRepo(dbPool)
	tenantService := service.NewTenantService(tenantRepo)
	tenantHandler := handler.NewTenantHandler(tenantService)

	router := routes.TenantRoutes(tenantHandler)

	server := &http.Server{
		Addr:         ":" + cfg.PORT,
		Handler:      apicommon.RecoverMiddleware(router),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	log.Printf("🚀 Tenant Service is running on port %s", cfg.PORT)
	if err := server.ListenAndServe(); err != nil {
		log.Printf("❌ Server error: %v", err)
		os.Exit(1)
	}
}
