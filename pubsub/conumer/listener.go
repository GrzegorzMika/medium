package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"time"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type dbNotification struct {
	Timestamp   time.Time `json:"timestamp"`
	SignalName  string    `json:"signal_name"`
	SignalValue float64   `json:"signal_value"`
}

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
	conn, err := l.aquireConnection(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	conn, err = l.listen(ctx, conn)
	if err != nil {
		return err
	}

	for {
		notification, err := conn.Conn().WaitForNotification(ctx)
		if err != nil {
			return fmt.Errorf("failed to wait for notification: %w", err)
		}
		l.handleNotifications(notification)
	}

}

func (l *Listener) aquireConnection(ctx context.Context) (*pgxpool.Conn, error) {
	conn, err := l.pool.Acquire(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to acquire connection from pool: %w", err)
	}
	return conn, nil
}

func (l *Listener) listen(ctx context.Context, conn *pgxpool.Conn) (*pgxpool.Conn, error) {
	_, err := conn.Exec(ctx, fmt.Sprintf("LISTEN %s", l.channelName))
	if err != nil {
		return nil, fmt.Errorf("failed to listen to channel: %w", err)
	}
	return conn, nil
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
