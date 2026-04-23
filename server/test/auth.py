import requests

from utils import GREEN, RED, RESET, SERVER_API, checkRespOk, deco, decoTitle

USERNAME = "USERNAME"
EMAIL = "EMAIL"
PASSWORD = "PASSWORD"

COOKIES = {}

@deco
def signup():
    resp = requests.post(SERVER_API+"/users", json={
        "username": USERNAME,
        "email": EMAIL,
        "password": PASSWORD,
    })
    checkRespOk(resp, "Signup1", 201)
    resp = requests.post(SERVER_API+"/users", json={
        "username": USERNAME,
        "email": EMAIL,
        "password": PASSWORD,
    })
    if resp.status_code != 409 or resp.json()["message"] != "username is already taken":
        print(f"{RED}failed test 'Signup2': {resp.status_code} {resp.content}{RESET}")
    else:
        print(f"{GREEN}Signup2: test passed{RESET}")

@deco
def postlogin():
    global COOKIES
    # success
    resp = requests.post(SERVER_API+"/sessions", json={
        "email": EMAIL,
        "password": PASSWORD,
    })
    sessionToken = resp.cookies.get("sessionToken")
    if sessionToken:
        checkRespOk(resp, "PostLogin1 valid", 201)
    else:
        checkRespOk(resp, "PostLogin1 valid", -1)
    COOKIES = resp.cookies.get_dict()
    # error wrong username
    resp = requests.post(SERVER_API+"/sessions", json={
        "email": "wrong",
        "password": PASSWORD,
    })
    checkRespOk(resp, "PostLogin2 error", 401)
    resp = requests.post(SERVER_API+"/sessions", json={
        "email": EMAIL,
        "password": "wrong",
    })
    checkRespOk(resp, "PostLogin3 error", 401)

@deco
def patchlogin():
    # valid
    resp = requests.patch(SERVER_API+"/sessions", cookies=COOKIES)
    checkRespOk(resp, "Login2", 200)
    # error
    resp = requests.patch(SERVER_API+"/sessions", cookies={"sessionToken":"something random"})
    checkRespOk(resp, "Login2", 401, "session expired")

@deco
def logout():
    # valid
    resp = requests.delete(SERVER_API+"/sessions", cookies=COOKIES)
    checkRespOk(resp, "Login2", 200)
    # error
    resp = requests.delete(SERVER_API+"/sessions", cookies=COOKIES)
    checkRespOk(resp, "Login2", 401, "unauthorized")

@decoTitle
def auth():
    signup()
    postlogin()
    patchlogin()
    logout()

if __name__ == "__main__":
    auth()

