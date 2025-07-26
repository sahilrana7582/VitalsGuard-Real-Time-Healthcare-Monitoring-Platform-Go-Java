package main

import (
	"github.com/sahilrana7582/vitals-guard/auth-service/config"
	"github.com/sahilrana7582/vitals-guard/auth-service/internal/db"
)

func main() {

	cfg := config.LoadAuthConfig(".env")

	_ = db.NewConnPool(cfg)
}
