package main

import (
	"context"
	"log"
	"math/rand/v2"
	"time"
)

const (
	cap  = time.Minute
	base = time.Second
)

type Effector func(context.Context, DataProvider) ([]Results, error)

func Retry(effector Effector) Effector {
	return func(ctx context.Context, provider DataProvider) ([]Results, error) {
		if ctx.Err() != nil {
			return nil, ctx.Err()
		}
		results, err := effector(ctx, provider)
		for backoff := base; err != nil && backoff <= cap; backoff <<= 1 {
			log.Println("waiting", backoff, "seconds")
			jitter := rand.Int64N((int64(backoff)))
			time.Sleep(base + time.Duration(jitter))
			results, err = effector(ctx, provider)
		}
		return results, err
	}
}
