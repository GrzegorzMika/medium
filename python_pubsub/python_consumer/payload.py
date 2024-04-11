import datetime

from pydantic import BaseModel


class Payload(BaseModel):
    timestamp: datetime.datetime
    signal_name: str
    signal_value: float
