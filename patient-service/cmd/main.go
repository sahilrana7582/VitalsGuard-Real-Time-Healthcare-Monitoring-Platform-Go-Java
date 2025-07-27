package main

import (
	"github.com/sahilrana7582/vitals-guard/patient-service/config"
	"github.com/sahilrana7582/vitals-guard/patient-service/internal/db"
)

func main() {

	patientConfig := config.NewPatientConfig(".env")

	_ = db.NewConnPool(patientConfig)

}
