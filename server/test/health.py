import requests

import subprocess, sys, datetime, time
from utils import GREEN, PROJECT_ROOT, RED, RESET, SERVER_API, checkRespOk, deco, decoTitle

@deco
def startup():
    startTime = datetime.datetime.now().timestamp()

    while True:
        out = subprocess.run(["make logs"], cwd=PROJECT_ROOT, shell=True, capture_output=True).stdout.decode()
        now = datetime.datetime.now().timestamp()
        startupTime = now - startTime
        if out.find("Server init successful if no error") >= 0:
            time.sleep(1)
            if startupTime < 1:
                print(f"{GREEN}Server already running{RESET}")
            else:
                print(f"{GREEN}Server startup successfull\nStartup time: {RESET}", startupTime)
            return
        if startupTime >= 120:
            print(f"{RED}Error starting up server{RESET}")
            print(out)
            sys.exit(1)
        time.sleep(5)

@deco
def apiHealth():
    print("SERVER_API: ", SERVER_API)
    resp = requests.get(SERVER_API + "/health")
    checkRespOk(resp, "healthTest", 200)

@decoTitle
def health():
    startup()
    apiHealth()

if __name__ == "__main__":
    health()
