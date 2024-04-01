package main

import (
	"encoding/json"
	"time"
)

type dbNotification struct {
	Timestamp   time.Time `json:"timestamp"`
	SignalName  string    `json:"signal_name"`
	SignalValue float64   `json:"signal_value"`
}

func (d *dbNotification) UnmarshalJSON(b []byte) error {
	var err error
	aux := struct {
		Timestamp   string  `json:"timestamp"`
		SignalName  string  `json:"signal_name"`
		SignalValue float64 `json:"signal_value"`
	}{}

	if err := json.Unmarshal(b, &aux); err != nil {
		return err
	}

	d.Timestamp, err = time.Parse(time.RFC3339Nano, aux.Timestamp)
	if err != nil {
		d.Timestamp, err = time.Parse(time.RFC3339, aux.Timestamp)
		if err != nil {
			return err
		}
	}
	d.SignalName = aux.SignalName
	d.SignalValue = aux.SignalValue
	return nil
}
