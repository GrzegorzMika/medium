FROM python:3.12-slim AS builder

RUN apt-get update \
    && apt-get install gcc -y \
    && apt install -y libpq-dev \
    && apt-get clean

ENV PYTHONUNBUFFERED 1
ENV PYTHONDONTWRITEBYTECODE 1

COPY ./requirements.txt /app/requirements.txt

RUN pip install --user -r /app/requirements.txt && rm -rf /root/.cache/pip

FROM python:3.12-slim AS app

COPY --from=builder /root/.local /root/.local

COPY . /app/

ENV PATH=/root/.local/bin:$PATH

WORKDIR /app

CMD python main.py