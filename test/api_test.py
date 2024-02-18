import requests
import json

if __name__ == "__main__":
    url = "http://localhost:8888/login"  # Corrected URL
    headers = {"Content-Type": "application/json; charset=utf-8"}
    data = {
        "username": "Martin",
        "password": "Martin"
    }

    response = requests.post(url, headers=headers, json=data, verify=False)  # Disable SSL verification
    print("Status Code:", response.status_code)
    print("JSON Response", response.json())