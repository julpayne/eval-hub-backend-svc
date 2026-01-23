package abstractions

import (
	"time"

	"github.ibm.com/julpayne/eval-hub-backend-svc/pkg/api"
)

type Storage interface {
	// This is used to identify the storage implementation in the logs and error messages
	GetDatasourceName() string

	Ping(timeout time.Duration) error

	// Evaluation job operations
	CreateEvaluationJob(evaluation *api.EvaluationJobConfig) error
	GetEvaluationJob(id string) (*api.EvaluationJobResource, error)
	GetEvaluationJobs(summary bool, limit int, offset int, statusFilter string) (*api.EvaluationJobResourceList, error)
	DeleteEvaluationJob(id string, hardDelete bool) error
	UpdateBenchmarkStatusForJob(id string, status api.BenchmarkStatus) error
	UpdateEvaluationJobStatus(id string, state api.EvaluationJobState) error

	// Collection operations
	CreateCollection(collection *api.CollectionResource) error
	GetCollection(id string, summary bool) (*api.CollectionResource, error)
	GetCollections(limit int, offset int) (*api.CollectionResourceList, error)
	UpdateCollection(collection *api.CollectionResource) error
	DeleteCollection(id string) error

	// Close the storage connection
	Close() error
}
