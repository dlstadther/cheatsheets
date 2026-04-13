VENV_NAME = .venv
VENV_PYTHON = $(VENV_NAME)/bin/python


# clean:
# 	rm -rf ${VENV_NAME}

venv:
	python -m venv ${VENV_NAME} && \
	${VENV_PYTHON} -m pip install zensical ghp-import pytest-playwright && \
	${VENV_NAME}/bin/playwright install chromium

serve:
	${VENV_NAME}/bin/zensical serve

build:
	${VENV_NAME}/bin/zensical build

test: build
	${VENV_PYTHON} -m http.server 8000 --directory site & \
	SERVER_PID=$$!; \
	sleep 1; \
	${VENV_NAME}/bin/pytest tests/ --base-url http://localhost:8000 -v; \
	TEST_EXIT=$$?; \
	kill $$SERVER_PID 2>/dev/null || true; \
	exit $$TEST_EXIT

gh-deploy:
	${VENV_NAME}/bin/zensical build && \
	${VENV_NAME}/bin/ghp-import -n -p -f site
