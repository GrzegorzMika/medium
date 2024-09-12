package main

import "time"

type Results struct {
	Timestamp  time.Time `json:"timestamp"`
	Prediction float32   `json:"prediction"`
}
