import requests
import json

if __name__ == "__main__":
    print("Login Test: Test for successful login with correct credentials.http://localhost:8888/login")
    url = "http://localhost:8888/login"  # Corrected URL
    headers = {"Content-Type": "application/json; charset=utf-8"}
    data = {
        "username": "Martin",
        "password": "Martin"
    }

    response = requests.post(url, headers=headers, json=data, verify=False)  # Disable SSL verification
    print("Status Code:", response.status_code)
    print("JSON Response", response.json())

    print("Register Test: Test for successful Registration with correct credentials.http://localhost:8888/register")
    url = "http://localhost:8888/register"  # Corrected URL
    headers = {"Content-Type": "application/json; charset=utf-8"}
    data = {
        "username": "Radek",
        "email":"radek@radek.com",
        "password": "Radek"
    }

    response = requests.post(url, headers=headers, json=data, verify=False)  # Disable SSL verification
    print("Status Code:", response.status_code)
    print("JSON Response", response.json())
