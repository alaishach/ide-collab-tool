import requests

from utils import SERVER_API, checkResp200

def healthCheck():
    resp = requests.get(SERVER_API + "/health")
    checkResp200(resp, "healtTest")
    print("Health Check Pass")

if __name__ == "__main__":
    healthCheck()
