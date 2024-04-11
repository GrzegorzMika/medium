CREATE TABLE IF NOT EXISTS signals (
    timestamp TIMESTAMPTZ NOT NULL,
    signal_name TEXT NOT NULL,
    signal_value FLOAT NOT NULL
);
