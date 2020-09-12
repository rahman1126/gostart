package dto

type HealthCheckResponse struct {
	Alive string `json:"alive"`
	Database string `json:"database"`
	Redis string `json:"redis"'`
}
