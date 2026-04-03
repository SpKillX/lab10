import pytest
import jwt
import requests
import datetime

GO_BASE_URL = "http://localhost:8080/api/v1/secure-resource"
SECRET_KEY = "university-course-work-secret-2026"

def test_go_auth_success():
    payload = {
        "sub": "test_user",
        "exp": datetime.datetime.now(datetime.UTC) + datetime.timedelta(minutes=15),
        "iat": datetime.datetime.now(datetime.UTC)
    }
    token = jwt.encode(payload, SECRET_KEY, algorithm="HS256")
    headers = {"Authorization": f"Bearer {token}"}
    
    response = requests.get(GO_BASE_URL, headers=headers)
    assert response.status_code == 200
    assert response.json()["status"] == "authorized"

def test_go_auth_no_header():
    response = requests.get(GO_BASE_URL)
    assert response.status_code == 401
    assert "error" in response.json()

def test_go_auth_wrong_secret():
    token = jwt.encode({"sub": "hacker"}, "WRONG_SECRET", algorithm="HS256")
    headers = {"Authorization": f"Bearer {token}"}
    
    response = requests.get(GO_BASE_URL, headers=headers)
    assert response.status_code == 401