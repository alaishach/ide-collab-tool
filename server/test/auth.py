import requests

from utils import GREEN, RED, SERVER_API, checkRespOk

USERNAME = "USERNAME"
EMAIL = "EMAIL"
PASSWORD = "PASSWORD"

def signup():
    resp = requests.post(SERVER_API+"/signup", json={
        "username": USERNAME,
        "email": EMAIL,
        "password": PASSWORD,
    })
    checkRespOk(resp, "Signup1", 201)
    resp = requests.post(SERVER_API+"/signup", json={
        "username": USERNAME,
        "email": EMAIL,
        "password": PASSWORD,
    })
    if resp.status_code != 409 or resp.json()["message"] != "username is already taken":
        print(f"{RED}failed test 'Signup2': {resp.status_code} {resp.content}")
    else:
        print(f"{GREEN}Signup2: test passed")

def login():
    resp = requests.post(SERVER_API+"/login", json={
        "username": USERNAME,
        "email": EMAIL,
        "password": PASSWORD,
    })
    if resp.status_code == 201:
        sessionToken = resp.cookies.get("sessionToken")
        if sessionToken:
            checkRespOk(resp, "Login1", 201)
        else:
            checkRespOk(resp, "Login1", -1)

def auth():
    signup()
    login()

if __name__ == "__main__":
    auth()

