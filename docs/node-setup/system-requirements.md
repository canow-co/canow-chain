# Hardware requirements

For most nodes, the RAM/vCPU requirements are relatively static and do not change over time. However, the disk storage space needs to grow as the chain grows and will evolve over time.

It is recommended to provide the disk storage as an expandable volume/partition that is mounted on your node configuration data path (the default is under `/home/canow`) so that it can be expanded independent of the root volume.

Extended information on [recommended hardware requirements is available in Tendermint documentation](https://docs.tendermint.com/main/tendermint-core/running-in-production.html#hardware). The figures below have been updated from the default Tendermint recommendations to account for current Canow network chain size, real-world usage accounting for requests nodes need to handle, etc.

## Minimum specifications

- 2 GB RAM
- x64 1.4 GHz 1 vCPU (or equivalent)
- 200 GB of disk space

## Recommended specifications

- 4 GB RAM
- x64 2.0 GHz 2 vCPU (or equivalent)
- 300 GB SSD

# Operating system

Our [packaged releases](https://github.com/canow-co/canow-chain/releases) are currently compiled and tested for `Ubuntu 22.04 LTS`, which is the recommended operating system for installation using interactive installer or binaries.

For other operating systems, we recommend using [pre-built Docker image releases for `canow-chain`](https://github.com/orgs/canow-co/packages?repo_name=canow-chain).

We plan on supporting other operating systems in the future, based on demand for specific platforms by the community.

# Storage volumes

We recommend using a storage path that can be kept persistent and restored/remounted (if necessary) for the configuration, data, and log directories associated with a node. This allows a node to be restored along with configuration files such as node keys and for the node's copy of the ledger to be restored without triggering a full chain sync.

The default directory location for `canow-chain` installations is `$HOME/.canow-chain`. Custom paths can be defined if desired.

# Ports

To function properly, `canow-chain` requires two types of ports to be configured. Depending on the setup, you may also need to configure firewall rules to allow the correct ingress/egress traffic.

Node operators should ensure there are no existing services running on these ports before proceeding with installation.

## P2P port

The P2P port is used for peer-to-peer communication between nodes. This port is used for your node to discover and connect to other nodes on the network. It should allow traffic to/from any IP address range.

- By default, the P2P port is set to `26656`.
- Inbound TCP connections on port `26656` (or your custom port) should be allowed from _any_ IP address range.
- Outbound TCP connections must be allowed on _all_ ports to _any_ IP address range.
- The default P2P port can be changed in `$HOME/.canow-chain/config/config.toml`.

Further details on [how P2P settings work is defined in Tendermint documentation](https://docs.tendermint.com/main/tendermint-core/running-in-production.html#p2p).

## RPC port

The RPC port is intended to be used by client applications as well as the canow-chain CLI. Your RPC port **must** be active and available on localhost to be able to use the CLI. It is up to a node operator whether they want to expose the RPC port to public internet.

The [RPC endpoints for a node](https://docs.tendermint.com/main/rpc/) provide REST, JSONRPC over HTTP, and JSONRPC over WebSockets. These API endpoints can provide useful information for node operators, such as healthchecks, network information, validator information etc.

- By default, the RPC port is set to `26657`
- Inbound and outbound TCP connections should be allowed from destinations desired by the node operator. The default is to allow this from any IPv4 address range.
- TLS for the RPC port can also be setup separately. Currently, TLS setup is not automatically carried out in the install process described below.
- The default RPC port can be changed in `$HOME/.canow-chain/config/config.toml`.

# Sentry nodes (optional)

Tendermint allows more complex setups in production, where the ingress/egress to a validator node is [proxied behind a "sentry" node](https://docs.tendermint.com/main/tendermint-core/validators.html).

While this setup is not compulsory, node operators with higher stakes or a need to have more robust network security may consider setting up a sentry-validator node architecture.
