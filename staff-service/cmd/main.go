package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sahilrana7582/vitals-guard/staff-service/config"
	"github.com/sahilrana7582/vitals-guard/staff-service/internal/db"
	"github.com/sahilrana7582/vitals-guard/staff-service/internal/handler"
	roleRepo "github.com/sahilrana7582/vitals-guard/staff-service/internal/repo/role"
	repo "github.com/sahilrana7582/vitals-guard/staff-service/internal/repo/staff"
	"github.com/sahilrana7582/vitals-guard/staff-service/internal/routes"
	roleService "github.com/sahilrana7582/vitals-guard/staff-service/internal/service/role"
	service "github.com/sahilrana7582/vitals-guard/staff-service/internal/service/staff"
)

func main() {
	cfg := config.NewStaffConfig(".env")

	connPool := db.NewConnPool(cfg)
	defer connPool.Close()

	roleRepo := roleRepo.NewRoleRepo(connPool)
	roleService := roleService.NewRoleService(roleRepo)
	roleHandler := handler.NewRoleHandler(roleService)

	staffRepo := repo.NewStaffRepo(connPool)
	staffService := service.NewStaffService(staffRepo)
	staffHandler := handler.NewStaffHandler(staffService)

	routes := routes.NewStaffRoutes(staffHandler, roleHandler)

	server := &http.Server{
		Addr:         ":" + cfg.PORT,
		Handler:      routes,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		log.Printf("🚀 Staff Service running on port %s", cfg.PORT)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("❌ Could not listen: %v\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("⚙️ Shutting down server gracefully...")
	if err := server.Close(); err != nil {
		log.Fatalf("❌ Server Shutdown Failed:%+v", err)
	}
	log.Println("✅ Server exited")
}
