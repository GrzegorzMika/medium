package main

import (
	"context"
	"errors"
	"log/slog"
	"sync"
	"time"
)

var ErrThrottleLimitReached = errors.New("throttle limit reached")

const (
	InitialTokensNumber = 2
	RefilRate           = 1
	RefillFrequency     = 10 * time.Second
)

type ResultsProviderFunc func(context.Context, int) ([]Results, error)

func Throttle(
	parentCtx context.Context,
	reader ResultsProviderFunc,
	initialTokensNumber int,
	refillRate int,
	refillFrequency time.Duration,
) ResultsProviderFunc {
	var tokens = initialTokensNumber
	var once sync.Once

	return func(ctx context.Context, length int) ([]Results, error) {
		if parentCtx.Err() != nil {
			return nil, parentCtx.Err()
		}
		if ctx.Err() != nil {
			return nil, ctx.Err()
		}
		once.Do(func() {
			ticker := time.NewTicker(refillFrequency)
			go func() {
				defer ticker.Stop()
				for {
					select {
					case <-parentCtx.Done():
						return
					case <-ticker.C:
						slog.Info("adding new token to the bucket")
						t := tokens + refillRate
						if t > initialTokensNumber {
							t = initialTokensNumber
						}
						tokens = t
					}
				}
			}()
		})

		if tokens <= 0 {
			return nil, ErrThrottleLimitReached
		}

		tokens--

		return reader(ctx, length)
	}
}
