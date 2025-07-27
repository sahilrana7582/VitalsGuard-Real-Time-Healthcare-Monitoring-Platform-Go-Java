package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sahilrana7582/vitals-guard/patient-service/config"
	"github.com/sahilrana7582/vitals-guard/patient-service/internal/db"
	"github.com/sahilrana7582/vitals-guard/patient-service/internal/handler"
	"github.com/sahilrana7582/vitals-guard/patient-service/internal/repo"
	"github.com/sahilrana7582/vitals-guard/patient-service/internal/routes"
	"github.com/sahilrana7582/vitals-guard/patient-service/internal/service"
)

func main() {

	patientConfig := config.NewPatientConfig(".env")

	connPool := db.NewConnPool(patientConfig)

	patientRepo := repo.NewPatientRepo(connPool)
	patientService := service.NewPatientService(patientRepo)
	patientHandler := handler.NewPatientHandler(patientService)

	patientRoutes := routes.NewPatientRoutes(patientHandler)

	srv := &http.Server{
		Addr:         ":" + patientConfig.PORT,
		Handler:      patientRoutes,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		log.Printf("🚀 Patient Service running on port %s", patientConfig.PORT)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("❌ Server failed: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("⚙️  Shutting down Patient Service...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("❌ Graceful shutdown failed: %v", err)
	}

	log.Println("✅ Patient Service stopped gracefully")

}
