# Setup local pool

/docker/localnet/docker-compose.env

```
BUILD_IMAGE="canow-co/canow-chain:build-latest"
```

Launch a local pool

* 4 validators
* 1 seed
* 1 observer

```
docker compose --env-file docker-compose.env up --force-recreate
```
