package main

import (
	"github.com/sahilrana7582/vitals-guard/tenant-service/config"
	"github.com/sahilrana7582/vitals-guard/tenant-service/internal/db"
)

func main() {
	cfg := config.LoadTenantsConfig(".env")

	_ = db.NewConnPool(cfg)

}
