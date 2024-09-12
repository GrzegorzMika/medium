package main

import (
	"context"
	"log"
)

func main() {
	ctx := context.Background()

	resultsProvider := NewResultsProvider()
	dataGetter := Retry(RequestData)

	results, err := dataGetter(ctx, resultsProvider)
	if err != nil {
		log.Println("failed to fetch data:", err)
		return
	}
	log.Printf("Fetched %d results", len(results))
}
