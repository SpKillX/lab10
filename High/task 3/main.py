import jwt
import datetime
import requests
from fastapi import FastAPI, HTTPException

app = FastAPI()
SECRET_KEY = "university-cou1rse-work-secret-2026"

@app.get("/request-secure-data")
async def get_data_from_go():
    payload = {
        "sub": "student_user",
        "exp": datetime.datetime.now(datetime.UTC) + datetime.timedelta(minutes=15),
        "iat": datetime.datetime.now(datetime.UTC)
    }
    
    token = jwt.encode(payload, SECRET_KEY, algorithm="HS256")
    
    headers = {"Authorization": f"Bearer {token}"}
    
    try:
        response = requests.get("http://localhost:8080/api/v1/secure-resource", headers=headers)
        return response.json()
    except requests.exceptions.RequestException as e:
        raise HTTPException(status_code=500, detail=f"Go service unreachable: {str(e)}")