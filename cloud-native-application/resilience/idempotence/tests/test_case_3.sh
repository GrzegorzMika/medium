curl -X "POST" -H "Content-Type: application/json" -d '{
    "signal_name": "testSensor",
    "slope": 0.25,
    "intercept": 2,
    "request_id": "e345e33a-0df3-400f-9f33-4ffa73d20285"
}' http://localhost:8080 &

curl -X "POST" -H "Content-Type: application/json" -d '{
    "signal_name": "testSensor",
    "slope": 0.25,
    "intercept": 2,
    "request_id": "e345e33a-0df3-400f-9f33-4ffa73d20285"
}' http://localhost:8080