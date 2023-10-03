# Canow chain installation

There are a few ways of canow-chain node installation:

- Binary installation (with/without [Cosmovisor](https://docs.cosmos.network/main/tooling/cosmovisor))
- Docker container setup

## Binary installation

```bash
# create a new directory and cd
mkdir ~/canow-chain-dir && cd ~/canow-chain-dir

# download the binary file for Linux x86-64
curl -L -o canow-chain.tar.gz https://github.com/canow-co/canow-chain/releases/download/v0.2.2/canow-chain-v0.2.2-linux.tar.gz

# unarchive the files
tar -xzvf canow-chain.tar.gz

# make canow-chain executable
chmod +x canow-chain

# install to the system
sudo cp canow-chain /usr/local/bin/canow-chain

# make sure that canow-chain works
canow-chain version
```

## Docker container

**This way is NOT recommended**. Especially for validator node because validator's key may be easily lost.

But if this way has been chosen, you can find pre-built image for all released versions in [the repo packages](https://github.com/canow-co/canow-chain/pkgs/container/canow-chain).

**Step1.** Pull the Docker image at the required `version`:

```bash
docker pull ghcr.io/canow-co/canow-chain:<version>
```

**Step2.** Start Docker image:

```bash
docker run -it -p "26657:26657" -p "26656:26656" -p "1317:1317" canow-chain
```

## Observer node configuration

### 1. Initialise the node configuration files

The "moniker" for your node is a "friendly" name that will be available to peers on the network in their address book, and allows easily searching a peer's address book.

```bash
canow-chain init <node-moniker>
```

### 2. Download the genesis file

Download the `genesis.json` file for the relevant [persistent chain](../../networks/) and put it in the `$HOME/.canow-chain/config` directory.

```bash
wget -O $HOME/.canow-chain/config/genesis.json https://raw.githubusercontent.com/canow-co/canow-chain/main/networks/testnet/genesis.json
```

### 3. Define the seed configuration for populating the list of peers known by a node

Update `seeds` with a comma separated list of seed node addresses specified in `seeds.txt` for the relevant [network](../../networks/).

For canow testnet, set the `SEEDS` environment variable:

```bash
SEEDS=$(wget -qO- https://raw.githubusercontent.com/canow-co/canow-chain/main/networks/testnet/seeds.txt)
```

After the `SEEDS` variable is defined, pass the values to the `canow-chain configure` tool to set it in the configuration file.

```bash
$ echo $SEEDS
# Comma separated list should be printed

$ canow-chain configure p2p seeds "$SEEDS"
```

### 4. Set gas prices accepted by the node

Update `minimum-gas-prices` parameter if you want to use custom value. The default is `25zarx`.

```bash
canow-chain configure min-gas-prices "25zarx"
```

### 5. Define the external peer-to-peer address

Unless you are running a node in a sentry/validator two-tier architecture, your node should be reachable on its peer-to-peer (P2P) port by other nodes. This can be defined by setting the `external-address` property which defines the externally reachable address. This can be defined using either IP address or DNS name followed by the P2P port (Default: 26656).

```bash
canow-chain configure p2p external-address <ip-address-or-dns-name:p2p-port>
# Example
# canow-chain configure p2p external-address 8.8.8.8:26656
```

This is especially important if the node has no public IP address, e.g., if it's in a private subnet with traffic routed via a load balancer or proxy. Without the `external-address` property, the node will report a private IP address from its own host network interface as its `remote_ip`, which will be unreachable from the outside world. The node still works in this configuration, but only with limited unidirectional connectivity.

### 6. Make the RPC endpoint available externally (optional)

This step is necessary only if you want to allow incoming client application connections to your node. Otherwise, the node will be accessible only locally.

```bash
canow-chain configure rpc-laddr "tcp://0.0.0.0:26657"
```

### 7. Enable and start the `canow-chain` system service

For setting up Canow node you can use cosmoviser for an automatically upgrade.

If you are prompted for a password for the `canow` user, type `exit` to logout and then attempt to execute this as a privileged user (with `sudo` privileges or as root user, if necessary).

```bash
  wget -O /etc/systemd/system/canow-chain.service https://github.com/canow-co/canow-chain/releases/download/v0.2.2/cosmovisor.service
  systemctl start canow-chain
```

Check that the `canow-chain` service is running. If successfully started, the status output should return `Active: active (running)`

```bash
systemctl status canow-chain
```

## Post-installation checks

Once the `canow-chain` daemon is active and running, check that the node is connected to the canow testnet and catching up with the latest updates on the ledger.

### Checking node status via terminal

```bash
canow-chain status
```

In the output, look for the text `latest_block_height` and note the value. Execute the status command above a few times and make sure the value of `latest_block_height` has increased each time.

The node is fully caught up when the parameter `catching_up` returns the output `false`.

### Checking node status via the RPC endpoint

An alternative method to check a node's status is via the RPC interface, if it has been configured.

- Remotely via the RPC interface: `canow-chain status --node <rpc-address>`
- By opening the JSONRPC over HTTP status page through a web browser: `<node-address:rpc-port>/status`

## Next steps

At this stage, your node would be connected to the canow testnet as an observer node. Learn [how to configure your node as a validator node](validator-creation.md) to participate in staking rewards, block creation, and governance.
