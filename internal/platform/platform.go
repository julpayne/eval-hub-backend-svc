package platform

import (
	"os"
	"strings"

	"github.com/julpayne/eval-hub-backend-svc/internal/constants"
)

func IsDevelopment() bool {
	return strings.EqualFold(os.Getenv(constants.EnvVarDevelopment), "true")
}
