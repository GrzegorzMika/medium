package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

const channelName = "new_signals"

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	pool, err := connectDB(ctx)
	if err != nil {
		panic(err)
	}

	listener := NewListener(pool, logCallback)
	err = listener.Start(ctx)
	if err != nil {
		panic(err)
	}
}

func connectDB(ctx context.Context) (*pgxpool.Pool, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	pool, err := pgxpool.New(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	return pool, nil
}

func logCallback(payload dbNotification) {
	slog.Info(fmt.Sprintf("notification received: %s-%s-%f", payload.Timestamp, payload.SignalName, payload.SignalValue))
}
