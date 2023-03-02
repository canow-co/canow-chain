# Canow chain installation

There are a few ways of canow-chain node installation:

* Binary installation (with/without [Cosmovisor](https://docs.cosmos.network/main/tooling/cosmovisor))
* Docker container setup

## Binary installation

## Docker container

**This way is NOT recommended**. Especially for validator node because validator's key may be easily lost. 

But if this way has been chosen, you can find pre-built image for all released versions in [the repo packages](https://github.com/canow-co/canow-chain/pkgs/container/canow-chain). 

**Step1.** Pull the Docker image at the required `version`:

```commandline
docker pull ghcr.io/canow-co/canow-chain:<version>
```


## Observer node configuration

