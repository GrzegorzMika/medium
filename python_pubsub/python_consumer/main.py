import logging
import os
from time import sleep

import psycopg

from listener import Listener
from payload import Payload

logging.basicConfig(level=logging.INFO)

CHANNEL_NAME = "new_signals"


def get_connection() -> psycopg.Connection:
    conn_string = os.getenv('DATABASE_URL')
    logging.info(f'Connecting to database: {conn_string}')

    while True:
        try:
            connection = psycopg.connect(conn_string)
            break
        except Exception:
            logging.info('Waiting for connection')
            sleep(1)

    return connection


def logging_callback(payload: Payload) -> None:
    logging.info('Received payload: %s', payload)


if __name__ == '__main__':
    conn = get_connection()

    listener = Listener(conn, CHANNEL_NAME)
    listener.add_callback(logging_callback).start()
