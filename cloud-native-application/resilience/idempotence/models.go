package main

import (
	"log/slog"
	"time"

	petname "github.com/dustinkirkland/golang-petname"
)

type DummyModelManager struct{}

func NewDummyModelManager() DummyModelManager {
	return DummyModelManager{}
}

func (m DummyModelManager) Create(signalName string, slope, intercept float64) (*Model, error) {
	time.Sleep(5 * time.Second)
	modelID := petname.Generate(3, "_")
	slog.With("signal_name", signalName).With("slope", slope).With("intercep", intercept).With("model_id", modelID).Info("Creating model")
	return &Model{
		ModelID:    modelID,
		SignalName: signalName,
		Slope:      slope,
		Intercept:  intercept,
	}, nil
}

func (m DummyModelManager) RollbackCreate(modelID string) error {
	slog.With("model_id", modelID).Info("Rolling back model creation")
	return nil
}
