# Docker

## Commands

Cleanup:
```shell
# remove all unused data, including volumes
docker system prune --volumes

# If wanting more granular cleanup...

# remove dangling build caches
docker builder prune
# remove stopped containers
docker container prune
# remove dangling images (those named `<none>`)
docker image prune
# remove unused networks
docker network prune
# remove anonymous local unused volumes
docker volume prune
```
