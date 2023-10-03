# Build and install Canow chain application

Setup Go according project [environment description](environment.md).

Pull repo from GitHub

```bash
git clone https://github.com/canow-co/canow-chain.git
cd canow-chain
```

## Build the project

```bash
make build
```

## Install binary

Move the binary to `/usr/bin/`

```bash
mv build/canow-chain /usr/bin/
```

## Define permissions (optional)

On this step we need to allow the binary execution if it is needed.

```bash
sudo chmod +x /usr/bin/canow-chain
```
