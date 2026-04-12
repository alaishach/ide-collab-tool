import dotenv, os, sys

from requests import Response

dotenv.load_dotenv()

def getEnv(key) -> str:
    value = os.getenv(key)
    if value is None:
        print(f"Could not find key: {key} inside env")
        sys.exit(1)
    return value

def checkResp200(resp: Response, testName: str):
    if resp.status_code != 200:
        print(f"failed test '{testName}': {resp.content}")
        sys.exit(1)

SERVER_API = getEnv("SERVER_API")
