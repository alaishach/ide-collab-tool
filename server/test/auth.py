import requests

from utils import SERVER_API, checkRespOk

def signup():
    resp = requests.post(SERVER_API+"/signup", json={
        "username": "werew",
        "email": "testEmail",
        "password": "testPassword",
    })
    checkRespOk(resp, "Signup")

def auth():
    signup()

if __name__ == "__main__":
    auth()

