# Setup local pool

/docker/localnet/build-latest.env

```bash
BUILD_IMAGE="canow-co/canow-chain:build-latest"
```

Launch a local pool

- 4 validators
- 1 seed
- 1 observer

```bash
docker compose --env-file build-latest.env up --force-recreate
```
