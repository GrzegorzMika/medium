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
	r := rand.New(rand.NewPCG(uint64(time.Now().UnixNano()), uint64(time.Now().UnixNano())))
	return func(ctx context.Context, provider DataProvider) ([]Results, error) {
		if ctx.Err() != nil {
			return nil, ctx.Err()
		}
		results, err := effector(ctx, provider)
		for backoff := base; err != nil && backoff <= cap; backoff <<= 1 {
			delay := base + time.Duration(r.Int64N((int64(backoff))))
			log.Println("waiting", delay, "seconds")
			time.Sleep(delay)
			results, err = effector(ctx, provider)
		}
		return results, err
	}
}
