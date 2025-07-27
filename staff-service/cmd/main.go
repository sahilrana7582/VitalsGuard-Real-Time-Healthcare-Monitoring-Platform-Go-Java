package main

import (
	"github.com/sahilrana7582/vitals-guard/staff-service/config"
	"github.com/sahilrana7582/vitals-guard/staff-service/internal/db"
)

func main() {
	cfg := config.NewStaffConfig(".env")

	_ = db.NewConnPool(cfg)

}
