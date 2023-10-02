# New validator creation

For setting up a new validator firstly [set up an observer node](canow-node-installation.md) and after this promote it to a validator.

To create your validator, just use the following command:

```commandline
canow-chain tx staking create-validator \
  --amount=1000000zarx \
  --pubkey=$(canow-chain tendermint show-validator) \
  --moniker="choose a moniker" \
  --chain-id=<chain_id> \
  --commission-rate="0.10" \
  --commission-max-rate="0.20" \
  --commission-max-change-rate="0.01" \
  --min-self-delegation="1000000" \
  --gas="auto" \
  --gas-prices="25zarx" \
  --from=<key_name>
```

Parameters required in the transaction above are:
* **amount** is an amount of tokens to stake
* **from** is a key alias of the node operator account that makes the initial stake
* **min-self-delegation** is a strictly positive integer that represents the minimum amount of self-delegated voting power your validator must always have. A min-self-delegation of 1000000000 means your validator will never have a self-delegation lower than 1ARX.
* **pubkey** is Node's bech32-encoded validator public key from the previous step
* **commission-rate** is Validator's commission rate
* **commission-max-rate** is Validator's maximum commission rate, expressed as a number with up to two decimal points. The value for this cannot be changed later.
* **commission-max-change-rate** is used to measure % point change over the commission-rate. E.g. 1% to 2% is a 100% rate increase, but only 1 percentage point. The value for this cannot be changed later.
* **chain-id**: Unique identifier for the chain.
  * For Canow Testnet, this is `canow-testnet-1`
* **gas** is a maximum gas to use for this specific transaction. Using auto uses Cosmos's auto-calculation mechanism, but can also be specified manually as an integer value.
* **gas-adjustment** (optional). If you're using auto gas calculation, this parameter multiplies the auto-calculated amount by the specified factor, e.g., 1.2. This is recommended so that it leaves enough margin of error to add a bit more gas to the transaction and ensure it successfully goes through.
* **gas-prices** is a minimum gas price set by the validator

## Status check

The status of your node can be checked in the network [Explorer](https://explorer.testnet.canowchain.com/desmos) or via [Omniflix](https://omniflix.testnet.canowchain.com) page for making delegations.