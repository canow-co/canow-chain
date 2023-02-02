#!/bin/bash
# shellcheck disable=SC2086

# Generates network configuration for an arbitrary amount of validators, observers, and seeds.
set -euo pipefail

# sed in MacOS requires extra argument
if [[ "$OSTYPE" == "darwin"* ]]; then
  SED_EXT='.orig'
else
  SED_EXT=''
fi

# Params
CHAIN_ID=${1:-"cheqd"} # First parameter, default is "cheqd"

VALIDATORS_COUNT=${2:-4} # Second parameter, default is 4
SEEDS_COUNT=${3:-1} # Third parameter, default is 1
OBSERVERS_COUNT=${4:-1} # Fourth parameter, default is 1

function init_node() {
  NODE_HOME=$1
  NODE_MONIKER=$2

  echo "[${NODE_MONIKER}] Initializing"

  canow-chain init "${NODE_MONIKER}" --chain-id "${CHAIN_ID}" --home "${NODE_HOME}" 2> /dev/null
  canow-chain tendermint show-node-id --home "${NODE_HOME}" > "${NODE_HOME}/node_id.txt"
  canow-chain tendermint show-validator --home "${NODE_HOME}" > "${NODE_HOME}/node_val_pubkey.txt"
}

function configure_node() {
  NODE_HOME=$1
  NODE_MONIKER=$2

  echo "[${NODE_MONIKER}] Configuring app.toml and config.toml"

  APP_TOML="${NODE_HOME}/config/app.toml"
  CONFIG_TOML="${NODE_HOME}/config/config.toml"

  sed -i $SED_EXT 's/minimum-gas-prices = ""/minimum-gas-prices = "50ncheq"/g' "${APP_TOML}"
  sed -i $SED_EXT 's/enable = false/enable = true/g' "${APP_TOML}"
  sed -i $SED_EXT 's|laddr = "tcp://127.0.0.1:26657"|laddr = "tcp://0.0.0.0:26657"|g' "${CONFIG_TOML}"
  sed -i $SED_EXT 's|addr_book_strict = true|addr_book_strict = false|g' "${CONFIG_TOML}"
  sed -i $SED_EXT 's/timeout_propose = "3s"/timeout_propose = "500ms"/g' "${CONFIG_TOML}"
  sed -i $SED_EXT 's/timeout_prevote = "1s"/timeout_prevote = "500ms"/g' "${CONFIG_TOML}"
  sed -i $SED_EXT 's/timeout_precommit = "1s"/timeout_precommit = "500ms"/g' "${CONFIG_TOML}"
  sed -i $SED_EXT 's/timeout_commit = "5s"/timeout_commit = "500ms"/g' "${CONFIG_TOML}"
  sed -i $SED_EXT 's/log_level = "info"/log_level = "debug"/g' "${CONFIG_TOML}"
}

function configure_genesis() {
  NODE_HOME=$1
  NODE_MONIKER=$2

  echo "[${NODE_MONIKER}] Configuring genesis"

  GENESIS="${NODE_HOME}/config/genesis.json"
  GENESIS_TMP="${NODE_HOME}/config/genesis_tmp.json"

  # Default denom
  sed -i $SED_EXT 's/"stake"/"ncheq"/' "${GENESIS}"

  # Short voting period
  sed -i $SED_EXT 's/"voting_period": "172800s"/"voting_period": "12s"/' "${GENESIS}"

  # Test accounts
  BASE_ACCOUNT_1="canow1xx3k0c86g9hclfh26u6gdjmdfkp75tvt3qa4mn"
  # Mnemonic: web region erupt kitchen ignore scout always cool advance tip window thank become liberty uncle reject powder task wheat industry blouse frozen trend two
  jq '.app_state.bank.balances += [{"address": "'${BASE_ACCOUNT_1}'", "coins": [{"denom": "ncheq", "amount": "100001000000000000"}] }]' "$GENESIS" > "$GENESIS_TMP" && \
    mv "${GENESIS_TMP}" "${GENESIS}"
  jq '.app_state.auth.accounts += [{"@type": "/cosmos.auth.v1beta1.BaseAccount","address": "'${BASE_ACCOUNT_1}'", "pub_key": null,"account_number": "0","sequence": "0"}]' "$GENESIS" > "$GENESIS_TMP" && \
    mv "${GENESIS_TMP}" "${GENESIS}"

  BASE_ACCOUNT_2="canow1p0xxlce6mvzh4kpdq06szr5z5uxrp9qx5gdt4k"
  # Mnemonic: field result budget animal friend solar update diesel sock almost casino play emotion pink honey conduct check witness copy eagle unlock genius brown dice
  jq '.app_state.bank.balances += [{"address": "'${BASE_ACCOUNT_2}'", "coins": [{"denom": "ncheq", "amount": "100001000000000000"}] }]' "$GENESIS" > "$GENESIS_TMP"  && \
    mv "${GENESIS_TMP}" "${GENESIS}"
  jq '.app_state.auth.accounts += [{"@type": "/cosmos.auth.v1beta1.BaseAccount","address": "'${BASE_ACCOUNT_2}'", "pub_key": null,"account_number": "0","sequence": "0"}]' "$GENESIS" > "$GENESIS_TMP" && \
    mv "${GENESIS_TMP}" "${GENESIS}"

  BASE_ACCOUNT_3="canow1m7qjmjc7lhm4ydly70mj6gsqc4pdynmzvprpxn"
  # Mnemonic: alien worry rent coil melt treat eager used pioneer truck warfare number glimpse describe pulse bar scout nurse twenty lab lunch defy blossom bridge
  jq '.app_state.bank.balances += [{"address": "'${BASE_ACCOUNT_3}'", "coins": [{"denom": "ncheq", "amount": "100"}] }]' "$GENESIS" > "$GENESIS_TMP"  && \
    mv "${GENESIS_TMP}" "${GENESIS}"
  jq '.app_state.auth.accounts += [{"@type": "/cosmos.auth.v1beta1.BaseAccount","address": "'${BASE_ACCOUNT_3}'", "pub_key": null,"account_number": "0","sequence": "0"}]' "$GENESIS" > "$GENESIS_TMP" && \
    mv "${GENESIS_TMP}" "${GENESIS}"

  BASE_ACCOUNT_4="canow16tf864x097ejh2wav793z4938j0lr2fg8l26rt"
  # Mnemonic: hope naive brief outdoor purchase abandon place output ten suffer grape cliff strike loud arch switch attract link panic retreat planet lion potato repair
  jq '.app_state.bank.balances += [{"address": "'${BASE_ACCOUNT_4}'", "coins": [{"denom": "ncheq", "amount": "100001000000000000"}] }]' "$GENESIS" > "$GENESIS_TMP"  && \
    mv "${GENESIS_TMP}" "${GENESIS}"
  jq '.app_state.auth.accounts += [{"@type": "/cosmos.auth.v1beta1.BaseAccount","address": "'${BASE_ACCOUNT_4}'", "pub_key": null,"account_number": "0","sequence": "0"}]' "$GENESIS" > "$GENESIS_TMP" && \
    mv "${GENESIS_TMP}" "${GENESIS}"

  BASE_ACCOUNT_5="canow1yd7n3hc2yh2gjtlxh6lwz3eqpz8k7uz0s0ncuj"
  # Mnemonic: later sentence pumpkin logic front area patch salmon insect quick topple hollow scissors purchase pluck focus climb food enforce private rotate abstract more advice
  jq '.app_state.bank.balances += [{"address": "'${BASE_ACCOUNT_5}'", "coins": [{"denom": "ncheq", "amount": "100001000000000000"}] }]' "$GENESIS" > "$GENESIS_TMP"  && \
    mv "${GENESIS_TMP}" "${GENESIS}"
  jq '.app_state.auth.accounts += [{"@type": "/cosmos.auth.v1beta1.BaseAccount","address": "'${BASE_ACCOUNT_5}'", "pub_key": null,"account_number": "0","sequence": "0"}]' "$GENESIS" > "$GENESIS_TMP" && \
    mv "${GENESIS_TMP}" "${GENESIS}"

  BASE_ACCOUNT_6="canow1h48utnqhmkzvlp76qv65da0ahtz64smz2yw2rg"
  # Mnemonic: margin burden miss kidney plug replace jaguar sound spoil notice lens inquiry laugh canvas firm sister fortune later tired asset scatter true athlete nice
  jq '.app_state.bank.balances += [{"address": "'${BASE_ACCOUNT_6}'", "coins": [{"denom": "ncheq", "amount": "100"}] }]' "$GENESIS" > "$GENESIS_TMP"  && \
    mv "${GENESIS_TMP}" "${GENESIS}"
  jq '.app_state.auth.accounts += [{"@type": "/cosmos.auth.v1beta1.BaseAccount","address": "'${BASE_ACCOUNT_6}'", "pub_key": null,"account_number": "0","sequence": "0"}]' "$GENESIS" > "$GENESIS_TMP" && \
    mv "${GENESIS_TMP}" "${GENESIS}"

  BASE_VESTING_ACCOUNT="canow1m6j32klalgrzpg6vlzhmkwtjj5aay9kn5ezl76"
  # Mnemonic: decide black crew connect cricket duck return finish piece license rough design lunch rude remember faculty shy cannon list toddler throw divide rent antique
  # shellcheck disable=SC2089
  BASE_VESTING_COIN="{\"denom\":\"ncheq\",\"amount\":\"10001000000000000\"}"
  jq '.app_state.bank.balances += [{"address": "'${BASE_VESTING_ACCOUNT}'", "coins": [{"denom": "ncheq", "amount": "5000000000000000"}] }]' "$GENESIS" > "$GENESIS_TMP" && \
    mv "${GENESIS_TMP}" "${GENESIS}"
  jq '.app_state.auth.accounts += [{"@type": "/cosmos.vesting.v1beta1.BaseVestingAccount", "base_account": {"address": "'${BASE_VESTING_ACCOUNT}'","pub_key": null,"account_number": "0","sequence": "0"}, "original_vesting": ['"${BASE_VESTING_COIN}"'], "delegated_free": [], "delegated_vesting": [], "end_time": "1630362459"}]' "$GENESIS" > "$GENESIS_TMP" && \
    mv "${GENESIS_TMP}" "${GENESIS}"

  CONTINUOUS_VESTING_ACCOUNT="canow1jnaaf6qyczz746cae67d79zg8cst2enuc6h9q7"
  # Mnemonic: pill soap false obvious echo still marine salute wheel patrol tourist sunset pizza destroy know alpha scare foot tragic lamp twin zero tonight defy
  jq '.app_state.bank.balances += [{"address": "'${CONTINUOUS_VESTING_ACCOUNT}'", "coins": [{"denom": "ncheq", "amount": "5000000000000000"}] }]' "$GENESIS" > "$GENESIS_TMP" && \
    mv "${GENESIS_TMP}" "${GENESIS}"
  jq '.app_state.auth.accounts += [{"@type": "/cosmos.vesting.v1beta1.ContinuousVestingAccount", "base_vesting_account": { "base_account": {"address": "'${CONTINUOUS_VESTING_ACCOUNT}'","pub_key": null,"account_number": "0","sequence": "0"}, "original_vesting": ['"${BASE_VESTING_COIN}"'], "delegated_free": [], "delegated_vesting": [], "end_time": "1630362459"}, "start_time": "1630352459"}]' "$GENESIS" > "$GENESIS_TMP" && \
    mv "${GENESIS_TMP}" "${GENESIS}"

  DELAYED_VESTING_ACCOUNT="canow197twl7px8chkezr4n4r9nmgw4mg4et90vpd7pc"
  # Mnemonic: grant sample panda define master just pink mesh trash bulk north nominee avocado car banner wide hip amateur boost seek basic ribbon phrase day
  jq '.app_state.bank.balances += [{"address": "'${DELAYED_VESTING_ACCOUNT}'", "coins": [{"denom": "ncheq", "amount": "5000000000000000"}] }]' "$GENESIS" > "$GENESIS_TMP" && \
    mv "${GENESIS_TMP}" "${GENESIS}"
  jq '.app_state.auth.accounts += [{"@type": "/cosmos.vesting.v1beta1.DelayedVestingAccount", "base_vesting_account": { "base_account": {"address": "'${DELAYED_VESTING_ACCOUNT}'","pub_key": null,"account_number": "0","sequence": "0"}, "original_vesting": ['"${BASE_VESTING_COIN}"'], "delegated_free": [], "delegated_vesting": [], "end_time": "1630362459"}}]' "$GENESIS" > "$GENESIS_TMP" && \
    mv "${GENESIS_TMP}" "${GENESIS}"

  PERIODIC_VESTING_ACCOUNT="canow1cxvfnmux4mfknpg0aya5hm649v076whwp60yvs"
  # Mnemonic: cattle deliver practice infant clip want tag exercise inch guilt equal license connect shoe boat high garage people slim party display demise lesson curious
  jq '.app_state.bank.balances += [{"address": "'${PERIODIC_VESTING_ACCOUNT}'", "coins": [{"denom": "ncheq", "amount": "5000000000000000"}] }]' "$GENESIS" > "$GENESIS_TMP" && \
    mv "${GENESIS_TMP}" "${GENESIS}"
  jq '.app_state.auth.accounts += [{"@type": "/cosmos.vesting.v1beta1.PeriodicVestingAccount", "base_vesting_account": { "base_account": {"address": "'${PERIODIC_VESTING_ACCOUNT}'","pub_key": null,"account_number": "0","sequence": "0"}, "original_vesting": ['"${BASE_VESTING_COIN}"'], "delegated_free": [], "delegated_vesting": [], "end_time": "1630362459"}, "start_time": "1630362439", "vesting_periods": [{"length": "20", "amount": ['"${BASE_VESTING_COIN}"']}]}]' "$GENESIS" > "$GENESIS_TMP" && \
    mv "${GENESIS_TMP}" "${GENESIS}"
}


NETWORK_CONFIG_DIR="network-config"
rm -rf $NETWORK_CONFIG_DIR
mkdir $NETWORK_CONFIG_DIR


# Generating node configurations
for ((i=0 ; i<VALIDATORS_COUNT ; i++))
do
  NODE_MONIKER="validator-$i"
  NODE_HOME="${NETWORK_CONFIG_DIR}/${NODE_MONIKER}"

  init_node "${NODE_HOME}" "${NODE_MONIKER}"
  configure_node "${NODE_HOME}" "${NODE_MONIKER}"
done

for ((i=0 ; i<SEEDS_COUNT ; i++))
do
  NODE_MONIKER="seed-$i"
  NODE_HOME="${NETWORK_CONFIG_DIR}/${NODE_MONIKER}"

  init_node "${NODE_HOME}" "${NODE_MONIKER}"
  configure_node "${NODE_HOME}" "${NODE_MONIKER}"
done

for ((i=0 ; i<OBSERVERS_COUNT ; i++))
do
  NODE_MONIKER="observer-$i"
  NODE_HOME="${NETWORK_CONFIG_DIR}/${NODE_MONIKER}"

  init_node "${NODE_HOME}" "${NODE_MONIKER}"
  configure_node "${NODE_HOME}" "${NODE_MONIKER}"
done


# Generating genesis
TMP_NODE_MONIKER="tmp"
TMP_NODE_HOME="${NETWORK_CONFIG_DIR}/${TMP_NODE_MONIKER}"
init_node "${TMP_NODE_HOME}" "${TMP_NODE_MONIKER}"
configure_genesis "${TMP_NODE_HOME}" "${TMP_NODE_MONIKER}"

mkdir "${TMP_NODE_HOME}/config/gentx"


# Adding genesis validators
for ((i=0 ; i<VALIDATORS_COUNT ; i++))
do
  NODE_MONIKER="validator-$i"
  NODE_HOME="${NETWORK_CONFIG_DIR}/${NODE_MONIKER}"

  cp "${TMP_NODE_HOME}/config/genesis.json" "${NODE_HOME}/config/genesis.json"

  canow-chain keys add "operator-$i" --keyring-backend "test" --home "${NODE_HOME}"
  canow-chain add-genesis-account "operator-$i" 20000000000000000ncheq --keyring-backend "test" --home "${NODE_HOME}"

  NODE_ID=$(canow-chain tendermint show-node-id --home "${NODE_HOME}")
  NODE_VAL_PUBKEY=$(canow-chain tendermint show-validator --home "${NODE_HOME}")
  canow-chain gentx "operator-$i" 1000000000000000ncheq --chain-id "${CHAIN_ID}" --node-id "${NODE_ID}" \
    --pubkey "${NODE_VAL_PUBKEY}" --keyring-backend "test"  --home "${NODE_HOME}"

  cp "${NODE_HOME}/config/genesis.json" "${TMP_NODE_HOME}/config/genesis.json"
  cp -R "${NODE_HOME}/config/gentx/." "${TMP_NODE_HOME}/config/gentx"
done


echo "Collecting gentxs"
canow-chain collect-gentxs --home "${TMP_NODE_HOME}"
canow-chain validate-genesis --home "${TMP_NODE_HOME}"

# Distribute final genesis
for ((i=0 ; i<VALIDATORS_COUNT ; i++))
do
  NODE_MONIKER="validator-$i"
  NODE_HOME="${NETWORK_CONFIG_DIR}/${NODE_MONIKER}"

  cp "${TMP_NODE_HOME}/config/genesis.json" "${NODE_HOME}/config/genesis.json"
done

for ((i=0 ; i<SEEDS_COUNT ; i++))
do
  NODE_MONIKER="seed-$i"
  NODE_HOME="${NETWORK_CONFIG_DIR}/${NODE_MONIKER}"

  cp "${TMP_NODE_HOME}/config/genesis.json" "${NODE_HOME}/config/genesis.json"
done

for ((i=0 ; i<OBSERVERS_COUNT ; i++))
do
  NODE_MONIKER="observer-$i"
  NODE_HOME="${NETWORK_CONFIG_DIR}/${NODE_MONIKER}"

  cp "${TMP_NODE_HOME}/config/genesis.json" "${NODE_HOME}/config/genesis.json"
done

# Leave one copy of genesis in the root of network-config
cp "${TMP_NODE_HOME}/config/genesis.json" "${NETWORK_CONFIG_DIR}/genesis.json"

# Generate seeds.txt
SEEDS_STR=""

for ((i=0 ; i<SEEDS_COUNT ; i++))
do
  NODE_MONIKER="seed-$i"
  NODE_P2P_PORT="26656"
  NODE_HOME="${NETWORK_CONFIG_DIR}/${NODE_MONIKER}"

  if ((i != 0))
  then
  SEEDS_STR="${SEEDS_STR},"
  fi

  SEEDS_STR="${SEEDS_STR}$(cat "${NODE_HOME}/node_id.txt")@${NODE_MONIKER}:${NODE_P2P_PORT}"
done

echo "${SEEDS_STR}" > "${NETWORK_CONFIG_DIR}/seeds.txt"

# We don't need the tmp node anymore
rm -rf "${TMP_NODE_HOME}"
