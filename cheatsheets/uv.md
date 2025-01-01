# UV

UV markets its self as "an extremely fast Python package and project manager, written in Rust".

As a project dependency manager and build tool, it's similar to Poetry.


[Docs](https://docs.astral.sh/uv)

Sample Conversion from Poetry to UV (0.3.2): [github.com/dlstadther/sample-project-python/pull/5](https://github.com/dlstadther/sample-project-python/pull/5)

Migration script to convert Poetry or Pipenv projects to UV: [mkniewallner/migrate-to-uv](https://github.com/mkniewallner/migrate-to-uv)


## Install
```shell
# latest
curl -LsSf https://astral.sh/uv/install.sh | sh

# specific version
curl -LsSf https://astral.sh/uv/0.3.0/install.sh | sh
```

## Upgrade
```shell
uv self update
```


## Installing Python
```shell
# view available and installed python versions
uv python list

# install latest patch version if major.minor version is not already installed
# (e.g. if 3.12.7 is installed, 3.12.8 won't be installed)
uv python install 3.12
# install a specific version
uv python install 3.12.8
```


## Use within Project
```shell
# initialize new project not yet created
uv init project-name

# initialize existing directory
cd project-name
uv init

# build project
uv build
# build project with single format
uv build sdist
uv build wheel

# publish project (assumes token set at UV_PUBLISH_TOKEN)
uv publish

# increment version (assuming Hatch is the build backend)
uvx hatch version patch
uvx hatch version minor
uvx hatch version major
uvx hatch version minor,rc
```
> `uv init` can be used within existing python projects, unless a `pyproject.toml` file has already been created using an alternative package manager (e.g. `poetry`)


## Manage Tools (outside of projects)
```shell
# temporary use
uvx ruff format

# installing tool for persistent access
uv tool install ruff
# use installed tool
ruff --version

# upgrade installed tool
uv tool upgrade ruff
```
