package main

import (
	"errors"
	"log/slog"
	"sync"
)

var ModelNotFoundError = errors.New("no model found for a given key")

type InMemoryStore struct {
	models map[string]*Model
	mu     *sync.RWMutex
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{
		models: make(map[string]*Model),
		mu:     &sync.RWMutex{},
	}
}

func (store *InMemoryStore) Get(key string) (*Model, error) {
	store.mu.RLock()
	defer store.mu.RUnlock()

	slog.With("key", key).Info("Checking if model exists in store")
	model, ok := store.models[key]
	if !ok {
		return nil, ModelNotFoundError
	}
	return model, nil
}

func (store *InMemoryStore) Set(key string, model *Model) error {
	store.mu.Lock()
	defer store.mu.Unlock()

	slog.With("key", key).Info("Saving model in store")
	store.models[key] = model
	return nil
}
