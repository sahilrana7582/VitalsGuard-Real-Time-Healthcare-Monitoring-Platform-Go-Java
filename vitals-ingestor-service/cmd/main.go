package main

import (
	"context"
	"log"
	"sync/atomic"
	"time"

	"github.com/sahilrana7582/vitals-guard/vitals-ingestor-service/config"
	"github.com/sahilrana7582/vitals-guard/vitals-ingestor-service/internal/db"
	"github.com/sahilrana7582/vitals-guard/vitals-ingestor-service/internal/models"
	"github.com/sahilrana7582/vitals-guard/vitals-ingestor-service/internal/pipeline"
	"github.com/sahilrana7582/vitals-guard/vitals-ingestor-service/internal/repo"
)

var vitalCounter uint64

func main() {
	cfg := config.NewConfig()

	connPool, err := db.NewDB(cfg)
	if err != nil {
		panic("Database Connection Refused")
	}

	vitalRepo := repo.NewRepo(connPool)

	vitalChan := make(chan *[]models.Vitals, 5000)

	go pipeline.StartFetcher(context.Background(), vitalRepo, vitalChan, 500, 30)

	go func() {
		ticker := time.NewTicker(2 * time.Second)
		defer ticker.Stop()

		for range ticker.C {
			count := atomic.SwapUint64(&vitalCounter, 0)
			log.Printf("\nProcessed %d vitals in the last 2 seconds\n", count)
		}
	}()

	for vitals := range vitalChan {
		go func(v *[]models.Vitals) {
			log.Printf("Fetched batch of %d vitals", len(*v))
			for _, _ = range *v {
				atomic.AddUint64(&vitalCounter, uint64(len(*v)))
			}
		}(vitals)
	}

}
