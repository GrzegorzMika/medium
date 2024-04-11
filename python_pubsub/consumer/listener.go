package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type callback func(dbNotification)

type Listener struct {
	pool        *pgxpool.Pool
	channelName string
	callbacks   []callback
}

func NewListener(pool *pgxpool.Pool, callbacks ...callback) *Listener {
	return &Listener{
		pool:        pool,
		channelName: channelName,
		callbacks:   callbacks,
	}
}

func (l *Listener) Start(ctx context.Context) error {
	slog.Info("Starting listener")

	// acquire a connection from the pool
	conn, err := l.pool.Acquire(ctx)
	if err != nil {
		return fmt.Errorf("failed to acquire connection from pool: %w", err)
	}
	// defer release the connection back to the pool
	defer conn.Release()

	// let database know we're ready to receive notifications
	_, err = conn.Exec(ctx, fmt.Sprintf("LISTEN %s", l.channelName))
	if err != nil {
		return fmt.Errorf("failed to listen to channel: %w", err)
	}

	// wait for notifications and process them
	for {
		notification, err := conn.Conn().WaitForNotification(ctx)
		if err != nil {
			return fmt.Errorf("failed to wait for notification: %w", err)
		}
		l.handleNotifications(notification)
	}

}

func (d *Listener) handleNotifications(notification *pgconn.Notification) {
	payload, err := d.parsePayload(notification.Payload)
	if err != nil {
		slog.With("error", err).Error("Could not parse PostgresSQL notification payload.")
	}
	for _, callback := range d.callbacks {
		callback(payload)
	}
}

func (d *Listener) parsePayload(rawPayload string) (dbNotification, error) {
	var payload dbNotification

	err := json.Unmarshal([]byte(rawPayload), &payload)
	if err != nil {
		return payload, fmt.Errorf("failed to unmarshal payload: %w", err)
	}

	return payload, nil
}
