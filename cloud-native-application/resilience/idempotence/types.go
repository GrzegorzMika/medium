package main

import "golang.org/x/sync/singleflight"

type CreateRequest struct {
	RequestID  string  `json:"request_id"`
	SignalName string  `json:"signal_name"`
	Slope      float64 `json:"slope"`
	Intercept  float64 `json:"intercept"`
}

type CreateResponse struct {
	ModelID string `json:"model_id"`
}

type Model struct {
	ModelID    string  `json:"model_id"`
	SignalName string  `json:"signal_name"`
	Slope      float64 `json:"slope"`
	Intercept  float64 `json:"intercept"`
}

type Cache interface {
	Get(key string) (*Model, error)
	Set(key string, model *Model) error
}

type ModelManager interface {
	Create(signalName string, slope, intercept float64) (*Model, error)
	RollbackCreate(modelID string) error
}

type App struct {
	cache   Cache
	group   singleflight.Group
	manager ModelManager
}
