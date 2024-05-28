package main

import (
	"context"
	"math/rand/v2"
	"time"
)

type Results struct {
	Timestamp  time.Time `json:"timestamp"`
	Prediction float32   `json:"prediction"`
}

type ResultsProvider struct{}

func NewResultsProvider() *ResultsProvider {
	return &ResultsProvider{}
}

func (p *ResultsProvider) GetResults(ctx context.Context, length int) ([]Results, error) {
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}
	time.Sleep(time.Duration(100*rand.Float32()) * time.Millisecond)
	results := make([]Results, 0, length)
	for i := range length {
		results = append(results, Results{
			Timestamp:  time.Now().Add(time.Duration(i) * time.Second),
			Prediction: rand.Float32(),
		})
	}
	return results, nil
}
