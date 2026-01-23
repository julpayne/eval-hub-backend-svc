package storage_sql

import (
	"context"
	"database/sql"
	"log/slog"
	"time"

	// import the postgres driver - "pgx"
	_ "github.com/jackc/pgx/v5/stdlib"
	// import the sqlite driver - "sqlite"
	_ "modernc.org/sqlite"

	"github.ibm.com/julpayne/eval-hub-backend-svc/internal/abstractions"
	"github.ibm.com/julpayne/eval-hub-backend-svc/internal/config"
	"github.ibm.com/julpayne/eval-hub-backend-svc/pkg/api"
)

const (
	DriverPostgres = "pgx"
	DriverSQLite   = "sqlite"
)

type SQLStorage struct {
	sqlConfig *config.SQLDatabaseConfig
	pool      *sql.DB
	ctx       context.Context
}

func NewSQLStorage(sqlConfig *config.SQLDatabaseConfig, logger *slog.Logger) (abstractions.Storage, error) {
	logger.Info("Creating SQL storage", "driver", sqlConfig.Driver, "url", sqlConfig.URL)

	pool, err := sql.Open(sqlConfig.Driver, sqlConfig.URL)
	if err != nil {
		return nil, err
	}

	if sqlConfig.ConnMaxLifetime != nil {
		pool.SetConnMaxLifetime(*sqlConfig.ConnMaxLifetime)
	}
	if sqlConfig.MaxIdleConns != nil {
		pool.SetMaxIdleConns(*sqlConfig.MaxIdleConns)
	}
	if sqlConfig.MaxOpenConns != nil {
		pool.SetMaxOpenConns(*sqlConfig.MaxOpenConns)
	}

	storage := &SQLStorage{
		sqlConfig: sqlConfig,
		pool:      pool,
		ctx:       context.Background(),
	}

	logger.Info("Pinging SQL storage", "driver", sqlConfig.Driver, "url", sqlConfig.URL)
	err = storage.Ping(1 * time.Second)
	if err != nil {
		return nil, err
	}

	return storage, nil
}

// Ping the database to verify DSN provided by the user is valid and the
// server accessible. If the ping fails exit the program with an error.
func (s *SQLStorage) Ping(timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(s.ctx, timeout)
	defer cancel()

	return s.pool.PingContext(ctx)
}

func (s *SQLStorage) GetDatasourceName() string {
	return s.sqlConfig.Driver
}

func (s *SQLStorage) CreateEvaluationJob(evaluation *api.EvaluationJobConfig) error {
	return nil
}

func (s *SQLStorage) GetEvaluationJob(id string) (*api.EvaluationJobResource, error) {
	return nil, nil
}

func (s *SQLStorage) GetEvaluationJobs(summary bool, limit int, offset int, statusFilter string) (*api.EvaluationJobResourceList, error) {
	return nil, nil
}

func (s *SQLStorage) DeleteEvaluationJob(id string, hardDelete bool) error {
	return nil
}

func (s *SQLStorage) UpdateBenchmarkStatusForJob(id string, status api.BenchmarkStatus) error {
	return nil
}

func (s *SQLStorage) UpdateEvaluationJobStatus(id string, state api.EvaluationJobState) error {
	return nil
}

func (s *SQLStorage) CreateCollection(collection *api.CollectionResource) error {
	return nil
}

func (s *SQLStorage) GetCollection(id string, summary bool) (*api.CollectionResource, error) {
	return nil, nil
}

func (s *SQLStorage) GetCollections(limit int, offset int) (*api.CollectionResourceList, error) {
	return nil, nil
}

func (s *SQLStorage) UpdateCollection(collection *api.CollectionResource) error {
	return nil
}

func (s *SQLStorage) DeleteCollection(id string) error {
	return nil
}

func (s *SQLStorage) Close() error {
	return s.pool.Close()
}
