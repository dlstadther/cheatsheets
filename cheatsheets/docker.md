# Docker

## Commands

Remove all dangling images (those named `<none>`):
```shell
docker rmi $(docker images -r "dangling=true" -q)
```
