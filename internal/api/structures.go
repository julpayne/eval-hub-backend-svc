package api

import "time"

// State represents the evaluation state enum
type State string

const (
	StatePending   State = "pending"
	StateRunning   State = "running"
	StateCompleted State = "completed"
	StateFailed    State = "failed"
	StateCancelled State = "cancelled"
)

// PatchOp represents the patch operation enum
type PatchOp string

// The tenant that provide scoping for objests stored in the database but not limited to the database.
type Tenant string

const (
	PatchOpReplace PatchOp = "replace"
	PatchOpAdd     PatchOp = "add"
	PatchOpRemove  PatchOp = "remove"
)

type Ref struct {
	ID string `json:"id"`
}

type HRef struct {
	Href string `json:"href"`
}

// Error represents an error response
type Error struct {
	Detail string `json:"detail"`
}

// PatchOperation represents a single patch operation
type PatchOperation struct {
	Op    PatchOp `json:"op"`
	Path  string  `json:"path"`
	Value any     `json:"value,omitempty"`
}

// Patch represents a list of patch operations
type Patch []PatchOperation

// Resource represents base resource fields
type Resource struct {
	ID        string    `json:"id"`
	Tenant    Tenant    `json:"tenant"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Page represents generic pagination schema
type Page struct {
	First      *HRef `json:"first"`
	Next       *HRef `json:"next,omitempty"`
	Limit      int   `json:"limit"`
	TotalCount int   `json:"total_count"`
}

// HealthResponse represents health check response
type HealthResponse struct {
	Status            string                    `json:"status"`
	Version           string                    `json:"version"`
	Timestamp         *time.Time                `json:"timestamp,omitempty"`
	Components        map[string]map[string]any `json:"components,omitempty"`
	UptimeSeconds     float64                   `json:"uptime_seconds"`
	ActiveEvaluations int                       `json:"active_evaluations,omitempty"`
}
