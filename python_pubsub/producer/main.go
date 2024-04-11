package main

import (
	"context"
	"fmt"
	"log/slog"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

const defaultReadHeaderTimeoutSeconds = 10 * time.Second

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	db, err := connectDB(ctx)
	if err != nil {
		slog.With("error", err).Error("error connecting to database")
		return
	}
	defer db.Close()

	mux := http.NewServeMux()

	mux.HandleFunc("POST /ingest", createIngestHandler(ctx, db))

	server := &http.Server{
		Addr:              ":8080",
		Handler:           mux,
		ReadHeaderTimeout: defaultReadHeaderTimeoutSeconds,
	}

	slog.Info("Start ingesting data")
	err = server.ListenAndServe()
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

func createIngestHandler(ctx context.Context, pool *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		signalName, signalValue := generateTestData()
		slog.With("signal name", signalName).With("signal value", signalValue).Info("Inserting data into database")
		_, err := pool.Exec(ctx, "INSERT INTO signals (timestamp, signal_name, signal_value) VALUES ($1, $2, $3)", time.Now().Format(time.RFC3339), signalName, signalValue)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(fmt.Sprintf("Failed to insert signal: %s", err)))
			return
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(fmt.Sprintf("Inserted signal %s with value %f", signalName, signalValue)))
	}
}

func generateTestData() (string, float64) {
	return "test_signal", rand.Float64()
}
