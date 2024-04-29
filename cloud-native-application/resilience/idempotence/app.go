package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

func (app *App) Run() error {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /", app.CreateModelHandler)

	server := &http.Server{
		Addr:              ":8080",
		Handler:           mux,
		ReadHeaderTimeout: 10 * time.Second,
	}

	return server.ListenAndServe()
}

func (app *App) CreateModelHandler(w http.ResponseWriter, r *http.Request) {
	var request CreateRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	model, err := app.cache.Get(request.RequestID)
	if err == nil && model != nil {
		slog.With("model_id", model.ModelID).Info("Get model from cache")
		_ = json.NewEncoder(w).Encode(CreateResponse{
			ModelID: model.ModelID,
		})
		return
	}

	modelRecord, err, _ := app.group.Do(request.RequestID, func() (any, error) {
		model, err = app.manager.Create(request.SignalName, request.Intercept, request.Slope)
		if err != nil {
			return nil, fmt.Errorf("failed to create model: %w", err)
		}

		err = app.cache.Set(request.RequestID, model)
		if err != nil {
			rollbackError := app.manager.RollbackCreate(model.ModelID)
			if rollbackError != nil {
				err = errors.Join(err, rollbackError)
			}
			return nil, fmt.Errorf("failed to save model to cache: %w", err)
		}
		return model, nil
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	model = modelRecord.(*Model)
	response := CreateResponse{
		ModelID: model.ModelID,
	}

	_ = json.NewEncoder(w).Encode(response)
}
