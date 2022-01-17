# Mkdocs

## Install
```shell
pip install mkdocs
```

## Default Project Structure
```
mkdocs.yml
docs/
- index.md
```

## Override Docs Directory
```yaml
# mkdocs.yml
docs_dir: cheatsheets
```

## Dev Server
```shell
# in same dir as mkdocs.yml
mkdocs serve

# open http://127.0.0.1:8000/
```

## Build the Site
```shell
mkdocs build
```

## Navigation Configuration
If specified, the navigation will present only that which is specified.
```yaml
# mkdocs.yml
nav:
    - Home: index.md
    - About: about.md
```
If no navigation is provided, defaults to all contents inside of the `docs_dir` in alphanumeric order.
