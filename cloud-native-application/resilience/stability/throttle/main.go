package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	provider := NewResultsProvider()
	mux := http.NewServeMux()

	mux.HandleFunc("GET /notthrottled", getNotThrottledHandler(ctx, provider))
	mux.HandleFunc("GET /throttled", getThrottledHandler(ctx, provider))

	httpServer := &http.Server{
		Addr:              ":8080",
		Handler:           mux,
		ReadHeaderTimeout: 10 * time.Second,
	}

	slog.With("address", httpServer.Addr).Info("starting HTTP server")
	err := httpServer.ListenAndServe()
	if err != nil {
		slog.With("error", err).Error("error serving requests")
	}
}

func getNotThrottledHandler(_ context.Context, provider *ResultsProvider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("serving not throttled request")
		length, err := parseLengthParameter(r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(err.Error()))
			return
		}

		results, err := provider.GetResults(r.Context(), length)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(err.Error()))
			return
		}

		err = json.NewEncoder(w).Encode(results)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(err.Error()))
			return
		}
	}
}

func getThrottledHandler(ctx context.Context, provider *ResultsProvider) http.HandlerFunc {
	reader := Throttle(ctx, provider.GetResults, InitialTokensNumber, RefilRate, RefillFrequency)
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("serving throttled request")
		length, err := parseLengthParameter(r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(err.Error()))
			return
		}

		results, err := reader(r.Context(), length)
		if err != nil && errors.Is(err, ErrThrottleLimitReached) {
			w.WriteHeader(http.StatusTooManyRequests)
			_, _ = w.Write([]byte(err.Error()))
			return
		}
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(err.Error()))
			return
		}

		err = json.NewEncoder(w).Encode(results)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(err.Error()))
			return
		}
	}
}

func parseLengthParameter(r *http.Request) (int, error) {
	length := r.URL.Query().Get("length")
	if length == "" {
		return 0, fmt.Errorf("length parameter is required")
	}
	lengthInt, err := strconv.Atoi(length)
	if err != nil {
		return 0, fmt.Errorf("failed to parse length parameter")
	}
	return lengthInt, nil
}
