package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.ibm.com/julpayne/eval-hub-backend-svc/internal/execution_context"
)

func (h *Handlers) HandleHealth(ctx *execution_context.ExecutionContext, w http.ResponseWriter) {
	if !h.checkMethod(ctx, http.MethodGet, w) {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":    "healthy",
		"timestamp": time.Now().UTC().Format(time.RFC3339),
	})
}
