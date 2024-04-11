from __future__ import annotations

import logging
from typing import Protocol

import psycopg

from payload import Payload


class Callback(Protocol):
    def __call__(self, payload: Payload) -> None:
        ...


class Listener:
    def __init__(self, connection: psycopg.Connection, channel_name: str):
        self.connection: psycopg.Connection = connection
        self.channel_name: str = channel_name
        self.callbacks: list[Callback] = []

    def add_callback(self, callbacks: Callback) -> Listener:
        self.callbacks.append(callbacks)
        return self

    def start(self) -> None:
        cursor = self.connection.cursor()
        cursor.execute(f"LISTEN {self.channel_name};")
        self.connection.commit()
        cursor.close()
        notifications = self.connection.notifies()
        for notification in notifications:
            payload = Payload.model_validate_json(notification.payload)
            for callback in self.callbacks:
                try:
                    callback(payload)
                except Exception as e:
                    logging.error(e)
