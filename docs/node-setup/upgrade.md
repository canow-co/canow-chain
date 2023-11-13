# Upgrade Procedure

1. Ensure that `canow-chain` client on your machine has the version `v0.2.1`:

   ```bash
   canow-chain version
   ```

2. Ensure that `node`, `chain-id` and `keyring-backend` client config parameters have proper values:

   ```bash
   canow-chain config node
   canow-chain config node https://observer.testnet.canowchain.com:26657

   canow-chain config chain-id
   canow-chain config chain-id canow-testnet-1

   canow-chain config keyring-backend
   canow-chain config keyring-backend os
   ```

3. Ensure that `min_deposit` equals to `8000000000000zarx`:

   ```bash
   canow-chain query params subspace gov depositparams
   ```

4. Ensure that `voting_period` equals to `172800000000000ns`:

   ```bash
   canow-chain query params subspace gov votingparams
   ```

5. Find out `latest_block_height`:

   ```bash
   canow-chain status
   ```

6. Taking into account that `testnet` network orders a new block every `5.88s`, calcualte `UPGRADE_HEIGHT` using the following formula:

   `UPGRADE_HEIGHT = latest_block_height + 29388 + 1224`

   where `29388` is a height increment for 2-day voting period (`172800000000000ns`) and `1224` is a height increment for 2-hour reserve after the voting end.

7. Submit the upgrade proposal:

   ```bash
   canow-chain tx gov submit-legacy-proposal software-upgrade v0.3.0 --title "Proposal of software upgrade to canow-chain v0.3.0" --description "This proposal is to upgrade canow-chain software installed on testnet network from [v0.2.1](https://github.com/canow-co/canow-chain/releases/tag/v0.2.1) to [v0.3.0](https://github.com/canow-co/canow-chain/releases/tag/v0.3.0)." --upgrade-height UPGRADE_HEIGHT --upgrade-info '{"binaries":{"linux/amd64":"https://github.com/canow-co/canow-chain/releases/download/v0.3.0/canow-chain-v0.3.0-linux.tar.gz?checksum=f9e92ea05f1f6c80004bdfc43a90301f0ff237b6cc5a05acdd4353a310ecbd11"}}' --deposit 8000000000000zarx --from validator_account_1 --gas auto --gas-adjustment 1.3 --gas-prices 50zarx
   ```

   where `UPGRADE_HEIGHT` is the upgrade height calculated above.

8. Discover the submitted proposal (it is the one with the highest ID in the list of all the proposals):

   ```bash
   canow-chain query gov proposals
   ```

9. Select the submitted proposal individually:

   ```bash
   canow-chain query gov proposal PROPOSAL_ID
   ```

   where `PROPOSAL_ID` is the submitted proposal ID discovered above.

10. For each `VALIDATOR_OPERATOR` from `["validator_account_1", "validator_account_2", "validator_account_3", "validator_account_4"]` vote for the proposal:

    ```bash
    canow-chain tx gov vote PROPOSAL_ID yes --from VALIDATOR_OPERATOR --gas auto --gas-adjustment 1.3 --gas-prices 50zarx
    ```

    where `PROPOSAL_ID` is the proposal ID and `VALIDATOR_OPERATOR` is the given validator operator.

11. Check the votes for the proposal:

    ```bash
    canow-chain query gov votes PROPOSAL_ID
    ```

12. Wait for `voting_end_time` shown in the proposal:

    ```bash
    canow-chain query gov proposal PROPOSAL_ID
    ```

    where `PROPOSAL_ID` is the proposal ID.

13. Ensure that `status` of the proposal is `PROPOSAL_STATUS_PASSED`:

    ```bash
    canow-chain query gov proposal PROPOSAL_ID
    ```

    where `PROPOSAL_ID` is the proposal ID.

14. Ensure that the upgrade has been scheduled:

    ```bash
    canow-chain query upgrade plan
    ```

15. Wait for `latest_block_height` reaches `UPGRADE_HEIGHT` and then until the upgrade procedure is completed (`catching_up` field must eventually be set back to `false`):

    ```bash
    canow-chain status
    ```

16. Ensure that the upgrade has been applied:

    ```bash
    canow-chain query upgrade applied v0.3.0
    ```

17. Update `canow-chain` client on your machine to the same version as `testnet` network has been updated to:

    ```bash
    sudo rm -rf /tmp/canow-chain-release

    sudo mkdir /tmp/canow-chain-release

    cd /tmp/canow-chain-release

    sudo curl -L -o canow-chain.tar.gz https://github.com/canow-co/canow-chain/releases/download/v0.3.0/canow-chain-v0.3.0-linux.tar.gz

    sudo tar -xzvf canow-chain.tar.gz --no-same-owner

    sudo chmod +x canow-chain

    sudo cp canow-chain /usr/local/bin/

    cd ~

    sudo rm -rf /tmp/canow-chain-release
    ```

18. Ensure that `canow-chain` client on your machine has now the version `v0.3.0`:

    ```bash
    canow-chain version
    ```
