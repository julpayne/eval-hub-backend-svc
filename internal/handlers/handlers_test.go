package handlers_test

import (
	"testing"

	"github.ibm.com/julpayne/eval-hub-backend-svc/internal/execution_context"
	"github.ibm.com/julpayne/eval-hub-backend-svc/internal/handlers"
)

func TestNew(t *testing.T) {
	h := handlers.New()
	if h == nil {
		t.Error("New() returned nil")
	}
}

func createExecutionContext(method string, uri string) *execution_context.ExecutionContext {
	return &execution_context.ExecutionContext{
		Method: method,
		URI:    uri,
	}
}
