package handlers

import (
	"fmt"
	"net/http"

	"github.ibm.com/julpayne/eval-hub-backend-svc/internal/execution_context"
)

type Handlers struct{}

func New() *Handlers {
	return &Handlers{}
}

func (h *Handlers) checkMethod(ctx *execution_context.ExecutionContext, method string, w http.ResponseWriter) bool {
	if ctx.Method != method {
		http.Error(w, fmt.Sprintf("Method %s not allowed, expecting %s", ctx.Method, method), http.StatusMethodNotAllowed)
		return false
	}
	return true
}
