package execution_context

import (
	"log/slog"
	"net/http"
	"time"

	"github.ibm.com/julpayne/eval-hub-backend-svc/internal/config"
	"github.ibm.com/julpayne/eval-hub-backend-svc/internal/logging"
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

// NewExecutionContext creates a new ExecutionContext with default values. This function
// is called at the route level before invoking evaluation-related handlers to set up
// request-scoped context.
//
// The function automatically:
//   - Enhances the logger with request-specific fields via logging.LoggerWithRequest
//   - Sets default timeout (60 minutes) and retry attempts (3)
//   - Initializes an empty metadata map
//
// This enables automatic request ID tracking (from X-Global-Transaction-Id header or
// auto-generated UUID) and structured logging with consistent request metadata.
//
// Parameters:
//   - r: The HTTP request to extract context from
//   - logger: The base logger to enhance with request fields
//   - serviceConfig: The service configuration to include in the context
//
// Returns:
//   - *ExecutionContext: A new execution context ready for use in handlers
func NewExecutionContext(r *http.Request, logger *slog.Logger, serviceConfig *config.Config) *ExecutionContext {
	// Enhance logger with request-specific fields
	enhancedLogger := logging.LoggerWithRequest(logger, r)

	return &ExecutionContext{
		Logger:         enhancedLogger,
		Config:         serviceConfig,
		TimeoutMinutes: 60,
		RetryAttempts:  3,
		Metadata:       make(map[string]interface{}),
	}
}
