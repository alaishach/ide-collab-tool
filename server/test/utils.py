from typing import Callable, Optional
import dotenv, os, sys

from requests import Response

CWD = os.getcwd()
PROJECT_ROOT = CWD[0:CWD.find("server")+6]

GREEN = "\033[32m"
RED = "\033[91m"
BLUE = "\033[34m"
RESET = "\033[0m"

dotenv.load_dotenv()

def deco(func: Callable):
    def wrapper():
        print(f"{BLUE}" + 15 * "- ")
        print(f"TEST: {RESET}", func.__name__.upper())
        func()
        print("")
    return wrapper

def decoTitle(func: Callable):
    def wrapper():
        print(f"{BLUE}" + 10 * "-" + func.__name__ + 10 * "-")
        func()
        print("")
    return wrapper

def getEnv(key) -> str:
    value = os.getenv(key)
    if value is None:
        print(f"Could not find key: {key} inside env")
        sys.exit(1)
    return value

def checkRespOk(resp: Response, testName: str, expected_code: int, expected_message: Optional[str] = None):
    if resp.status_code != expected_code or (expected_message and resp.json()["message"] != expected_message):
        print(f"{RED}failed test '{testName}': {resp.status_code} {resp.content}")
        sys.exit(1)
    else:
        print(f"{GREEN}{testName}: test passed")

SERVER_API = getEnv("SERVER_API")
