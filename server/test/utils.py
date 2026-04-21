import dotenv, os, sys

from requests import Response

CWD = os.getcwd()
PROJECT_ROOT = CWD[0:CWD.find("server")+6]

GREEN = "\033[92m"
RED = "\033[91m"
RESET = "\033[0m"

dotenv.load_dotenv()

def getEnv(key) -> str:
    value = os.getenv(key)
    if value is None:
        print(f"Could not find key: {key} inside env")
        sys.exit(1)
    return value

def checkRespOk(resp: Response, testName: str):
    if resp.status_code < 200 or resp.status_code > 299:
        print(f"{RED}failed test '{testName}': {resp.content}")
        sys.exit(1)
    else:
        print(f"{GREEN}{testName}: test passed")

SERVER_API = getEnv("SERVER_API")
