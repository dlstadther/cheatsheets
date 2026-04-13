VENV_NAME = .venv
VENV_PYTHON = $(VENV_NAME)/bin/python


# clean:
# 	rm -rf ${VENV_NAME}

venv:
	python -m venv ${VENV_NAME} && \
	${VENV_PYTHON} -m pip install zensical ghp-import

serve:
	${VENV_NAME}/bin/zensical serve

build:
	${VENV_NAME}/bin/zensical build

gh-deploy:
	${VENV_NAME}/bin/zensical build && \
	${VENV_NAME}/bin/ghp-import -n -p -f site
