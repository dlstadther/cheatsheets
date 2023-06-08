# Docker

## Commands

Remove all dangling images (those named `<none>`):
```shell
docker rmi $(docker images -f "dangling=true" -q)
```
