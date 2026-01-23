package execution_context

import (
	"log/slog"
	"time"

	"github.ibm.com/julpayne/eval-hub-backend-svc/internal/config"
)

// ExecutionContext contains execution context for API operations. This pattern enables
// type-safe passing of configuration and state to evaluation-related handlers, which
// receive an ExecutionContext instead of a raw http.Request.
//
// The ExecutionContext contains:
//   - Logger: A request-scoped logger with enriched fields (request_id, method, uri, etc.)
//   - Config: The service configuration
//   - Evaluation-specific state: model info, timeouts, retries, metadata
type ExecutionContext struct {
	Logger       *slog.Logger
	Config       *config.Config
	Method       string
	URI          string
	RawQuery     string
	EvaluationID string
	ModelURL     string
	ModelName    string
	//BackendSpec    BackendSpec
	//BenchmarkSpec  BenchmarkSpec
	TimeoutMinutes int
	RetryAttempts  int
	StartedAt      *time.Time
	Metadata       map[string]interface{}
	MLflowClient   interface{}
	ExperimentName *string
}
