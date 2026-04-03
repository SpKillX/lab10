import pytest
from fastapi.testclient import TestClient
from main import app

client = TestClient(app)

def test_proxy_booking_flow(requests_mock):
    go_url = "http://localhost:8080/book"
    payload = {
        "user_id": 123,
        "table_id": 5,
        "email": "student@university.edu"
    }
    requests_mock.post(go_url, json={"message": "Booking confirmed"}, status_code=200)

    response = client.post("/remote-book", json=payload)
    
    assert response.status_code == 200
    assert response.json()["message"] == "Booking confirmed"