package handlers

import (
	"encoding/json"
	"net/http"

	"github.ibm.com/julpayne/eval-hub-backend-svc/internal/execution_context"
)

// HandleGetSystemMetrics handles GET /api/v1/metrics/system
func (h *Handlers) HandleGetSystemMetrics(ctx *execution_context.ExecutionContext, w http.ResponseWriter) {
	if !h.checkMethod(ctx, http.MethodGet, w) {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "System metrics not yet implemented",
	})
}
