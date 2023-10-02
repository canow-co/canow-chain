# Binary Installation

Canow-chain binary can be downloaded from the repo [Releases on GitHub](https://github.com/canow-co/canow-chain/releases). The last stable and recommended release is marked by the label `latest`.
The binary can be used like a CLI for node requests, transactions building and sending.

Besides `canow-chain` binary is complete application for launching Canow node.

### Step 1. Download

For installing `canow-chain` lets download and unpack it to `/usr/bin/` directory. If you are executing this command from the user without root permission, `sudo` will be needed for getting permissions for `/usr/bin/` directory.
Here is the command to install version 0.2.2. You can check which version is the latest now and download it [from the repository](https://github.com/canow-co/canow-chain/releases/latest).

#### For Linux

```commandline
sudo wget -c https://github.com/canow-co/canow-chain/releases/download/v0.2.2/canow-chain-v0.2.2-linux.tar.gz -O - | sudo tar -xz -C /usr/bin/
```

#### For macOS

```commandline
sudo wget -c https://github.com/canow-co/canow-chain/releases/download/v0.2.2/canow-chain-v0.2.2-darwin.tar.gz -O - | sudo tar -xz -C /usr/bin/
```

### Step 2. Define permissions

After downloading and unpack Canow-chain binary it may have `-rw-r--r--` permissions. On this step we need to allow the binary execution.

```commandline
sudo chmod +x /usr/bin/canow-chain
```

### Step 3. Check

For checking canow-chain installation success you can get a version of the installed binary:

```commandline
canow-chain version

> v0.2.2
```
