import os
import logging
from time import sleep
import psycopg

CHANNEL_NAME = 'new_signals'
logging.basicConfig(level=logging.INFO)


def get_connection():
    conn_string = os.getenv('DATABASE_URL')
    logging.info(f'Connecting to database: {conn_string}')

    while True:
        try:
            conn = psycopg.connect(conn_string)
            break
        except Exception:
            logging.info('Waiting for connection')
            sleep(1)

    return conn
    

if __name__ == '__main__':

    conn = get_connection()

    cursor = conn.cursor()
    cursor.execute(f"LISTEN {CHANNEL_NAME};")
    conn.commit()
    cursor.close()
    gen = conn.notifies()
    for notify in gen:
        logging.info(notify)
