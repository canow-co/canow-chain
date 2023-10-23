#!/bin/bash

set -euox pipefail

DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)"

pushd "$DIR/../../../../docker/localnet"

# Generate configs (make sure old binary is installed locally)
bash gen-network-config.sh
sudo chown -R 1000:1000 network-config

# Import keys
bash import-keys.sh

# Start network
docker compose --env-file testnet-latest.env up --detach --no-build

# TODO: Get rid of this sleep.
sleep 5

# Copy keys
sudo docker compose --env-file testnet-latest.env cp network-config/validator-0/keyring-test validator-0:/home/canow/.canow-chain
sudo docker compose --env-file testnet-latest.env cp network-config/validator-1/keyring-test validator-1:/home/canow/.canow-chain
sudo docker compose --env-file testnet-latest.env cp network-config/validator-2/keyring-test validator-2:/home/canow/.canow-chain
sudo docker compose --env-file testnet-latest.env cp network-config/validator-3/keyring-test validator-3:/home/canow/.canow-chain

# Restore permissions
sudo docker compose --env-file testnet-latest.env exec --user root validator-0 chown -R canow:canow /home/canow
sudo docker compose --env-file testnet-latest.env exec --user root validator-1 chown -R canow:canow /home/canow
sudo docker compose --env-file testnet-latest.env exec --user root validator-2 chown -R canow:canow /home/canow
sudo docker compose --env-file testnet-latest.env exec --user root validator-3 chown -R canow:canow /home/canow

popd
