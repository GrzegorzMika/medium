package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

const defaultReadHeaderTimeoutSeconds = 10

func defineReadHeaderTimeout() time.Duration {
	httpHeaderTimeout := os.Getenv("HTTP_HEADER_TIMEOUT")
	if httpHeaderTimeout == "" {
		return defaultReadHeaderTimeoutSeconds * time.Second
	}
	timeout, _ := strconv.Atoi(httpHeaderTimeout)
	return time.Duration(timeout) * time.Second
}

func main() {

	timeout := defineReadHeaderTimeout()

	mux := http.NewServeMux()

	mux.HandleFunc("GET /timeout", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(fmt.Sprintf("Timeout: %d", timeout)))
	})

	server := &http.Server{
		Addr:              ":8080",
		Handler:           mux,
		ReadHeaderTimeout: timeout,
	}

	log.Fatal(server.ListenAndServe())
}
