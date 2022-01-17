VENV_NAME = .venv
VENV_PYTHON = $(VENV_NAME)/bin/python


# clean:
# 	rm -rf ${VENV_NAME}

venv:
	python -m venv ${VENV_NAME} && \
	${VENV_PYTHON} -m pip install mkdocs mkdocs-material

serve:
	${VENV_PYTHON} -m mkdocs serve

build:
	${VENV_PTYHON} -m mkdocs build

gh-deploy:
	${VENV_PYTHON} -m mkdocs gh-deploy --force
