package main

import (
	"context"
	"fmt"
	"time"
)

type DataProvider interface {
	GetResults(ctx context.Context, length int) ([]Results, error)
}

func RequestData(ctx context.Context, provider DataProvider) ([]Results, error) {
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}
	if time.Now().Second()%5 != 0 {
		return nil, fmt.Errorf("failed to fetch data")
	}
	return provider.GetResults(ctx, 1000)
}
