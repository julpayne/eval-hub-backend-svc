package api

// SupportedBenchmark represents simplified benchmark reference for provider list
type SupportedBenchmark struct {
	ID string `json:"id"`
}

// Provider represents provider specification
type Provider struct {
	ID                  string               `json:"id"`
	Label               string               `json:"label"`
	SupportedBenchmarks []SupportedBenchmark `json:"supported_benchmarks,omitempty"`
}

// ProviderList represents response for listing providers
type ProviderList struct {
	TotalCount int        `json:"total_count"`
	Items      []Provider `json:"items"`
}
