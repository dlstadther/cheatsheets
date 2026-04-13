import subprocess
import time

import pytest


@pytest.fixture(scope="session", autouse=True)
def static_server():
    proc = subprocess.Popen(
        ["python", "-m", "http.server", "8000", "--directory", "site"],
        stdout=subprocess.DEVNULL,
        stderr=subprocess.DEVNULL,
    )
    time.sleep(1)
    yield
    proc.terminate()
    proc.wait()
