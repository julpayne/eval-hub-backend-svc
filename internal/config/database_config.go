package config

import "time"

type DatabaseType string

const (
	DatabaseTypeSQL DatabaseType = "sql"
)

type DatabaseConfig struct {
	SQL   map[string]SQLDatabaseConfig   `mapstructure:"sql,omitempty"`
	JSON  map[string]JSONDatabaseConfig  `mapstructure:"json,omitempty"`
	Other map[string]OtherDatabaseConfig `mapstructure:"other,omitempty"`
}

type SQLDatabaseConfig struct {
	Enabled         bool           `mapstructure:"enabled,omitempty"`
	Driver          string         `mapstructure:"driver"`
	URL             string         `mapstructure:"url"`
	ConnMaxLifetime *time.Duration `mapstructure:"conn_max_lifetime,omitempty"`
	MaxIdleConns    *int           `mapstructure:"max_idle_conns,omitempty"`
	MaxOpenConns    *int           `mapstructure:"max_open_conns,omitempty"`
	Fallback        bool           `mapstructure:"fallback,omitempty"`
	// Other map[string]any `mapstructure:",remain"`
}

type JSONDatabaseConfig struct {
	Enabled bool   `mapstructure:"enabled,omitempty"`
	Driver  string `mapstructure:"driver"`
	URL     string `mapstructure:"url"`
}

type OtherDatabaseConfig struct {
	Enabled bool   `mapstructure:"enabled,omitempty"`
	URL     string `mapstructure:"url"`
}
