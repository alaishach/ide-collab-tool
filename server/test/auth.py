import requests

from utils import GREEN, RED, SERVER_API, checkRespOk

def signup():
    resp = requests.post(SERVER_API+"/signup", json={
        "username": "eire",
        "email": "te",
        "password": "testPassword",
    })
    checkRespOk(resp, "Signup", 201)
    resp = requests.post(SERVER_API+"/signup", json={
        "username": "eire",
        "email": "te",
        "password": "testPassword",
    })
    if resp.status_code != 409 or resp.json()["message"] != "username is already taken":
        print(f"{RED}failed test 'Signup2': {resp.status_code} {resp.content}")
    else:
        print(f"{GREEN}Signup2: test passed")

def auth():
    signup()

if __name__ == "__main__":
    auth()

