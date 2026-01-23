package config

import (
	"github.ibm.com/julpayne/eval-hub-backend-svc/internal/abstractions"
)

type Config struct {
	Service  *ServiceConfig       `json:"service"`
	Database *DatabaseConfig      `json:"database"`
	Storage  abstractions.Storage `json:"-"`
}
