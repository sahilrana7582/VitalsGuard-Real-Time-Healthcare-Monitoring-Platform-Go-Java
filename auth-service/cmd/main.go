package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/sahilrana7582/vitals-guard/auth-service/config"
	"github.com/sahilrana7582/vitals-guard/auth-service/internal/db"
	"github.com/sahilrana7582/vitals-guard/auth-service/internal/handler"
	"github.com/sahilrana7582/vitals-guard/auth-service/internal/repo"
	"github.com/sahilrana7582/vitals-guard/auth-service/internal/routes"
	"github.com/sahilrana7582/vitals-guard/auth-service/internal/service"
)

func main() {
	log.SetPrefix("[AUTH-SERVICE] ")

	cfg := config.LoadAuthConfig(".env")

	pool := db.NewConnPool(cfg)
	defer pool.Close()

	authRepo := repo.NewAuthRepo(pool)
	authService := service.NewAuthService(authRepo)
	authHandler := handler.NewAuthHandler(authService)

	authRoutes := routes.NewAuthRoutes(authHandler)

	srv := &http.Server{
		Addr:         ":" + cfg.PORT,
		Handler:      authRoutes,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		log.Printf("✅ Auth service is running on port %s\n", cfg.PORT)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("❌ Server failed: %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	<-stop
	log.Println("🛑 Shutting down auth service...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("❌ Graceful shutdown failed: %v", err)
	}

	log.Println("✅ Auth service stopped cleanly.")
}
