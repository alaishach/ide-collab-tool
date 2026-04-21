import requests

import subprocess, sys, datetime, time
from utils import PROJECT_ROOT, SERVER_API, checkRespOk

def startup():
    startTime = datetime.datetime.now().timestamp()

    while True:
        out = subprocess.run(["make logs"], cwd=PROJECT_ROOT, shell=True, capture_output=True).stdout.decode()
        now = datetime.datetime.now().timestamp()
        startupTime = now - startTime
        if out.find("[GIN-debug] Listening and serving HTTP on :") >= 0:
            if startupTime < 1:
                print("Gin Server already running")
            else:
                print("Gin Server startup successfull\nStartup time: ", startupTime)
            return
        if startupTime >= 120:
            print("Error starting up server")
            print(out)
            sys.exit(1)
        time.sleep(5)

def healthCheck():
    startup()
    resp = requests.get(SERVER_API + "/health")
    checkRespOk(resp, "healtTest")
    print("Health Check Pass")

if __name__ == "__main__":
    healthCheck()
