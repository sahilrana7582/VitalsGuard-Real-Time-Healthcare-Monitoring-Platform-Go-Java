package pipeline

import (
	"context"
	"log"
	"time"

	"github.com/sahilrana7582/vitals-guard/vitals-ingestor-service/internal/models"
	"github.com/sahilrana7582/vitals-guard/vitals-ingestor-service/internal/repo"
)

func StartFetcher(ctx context.Context, repo *repo.Repo, vitalChan chan<- *[]models.Vitals, limit int, fetcherSize int) {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			log.Println("Vitals fetcher shutting down...")
			return

		case <-ticker.C:
			for i := 0; i < fetcherSize; i++ {
				offset := i * limit
				go fetch(ctx, repo, offset, limit, vitalChan)
			}
		}
	}
}

func fetch(ctx context.Context, repo *repo.Repo, offset, limit int, vitalChan chan<- *[]models.Vitals) {
	vitals, err := repo.FetchVitals(ctx, offset, limit)
	if err != nil {
		log.Printf("Fetcher offset=%d failed: %v", offset, err)
		return
	}

	if len(*vitals) > 0 {
		log.Printf("Fetched %d vitals at offset %d", len(*vitals), offset)
		vitalChan <- vitals
	}
}
