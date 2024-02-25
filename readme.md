# wasmapp
**wasmapp** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).

Welcome to the **`wasmapp`** repository, a blockchain application scaffolded with Ignite CLI and enhanced with the Ignite app for Wasm integration. This repository showcases a Cosmos-SDK blockchain project, version 0.50.4, integrating Wasm for smart contracts functionality.

## **Getting Started**

### **Prerequisites**

- Go 1.21.6
- Node (latest version)

### **Installation**

### Install Node

Install the current version of the node binary. Clone the repository and install the **`wasmapp`**:

```bash
git clone https://github.com/toschdev/wasmapp wasmapp
cd wasmapp
make install
```

### **Configure Node for wasmapp Testnet**

### Initialize Node

Replace **`YOUR_MONIKER`** with your own moniker:

```bash
wasmappd init YOUR_MONIKER --chain-id wasmapp
```

### Download Genesis

```bash
wget -O genesis.json https://raw.githubusercontent.com/toschdev/wasmapp/main/network/genesis.json
mv genesis.json ~/.wasmapp/config
```

### Configure Seed Node

Use a seed node to bootstrap:

```bash
sed -i 's/seeds = ""/seeds = "a33581756fdcdfc15715090436084431894b7f0a@65.109.102.149:26656"/' ~/.wasmapp/config/config.toml
```

### **Launch Node**

### Create Service File

Create a **`wasmapp.service`** file in the **`/etc/systemd/system`** folder. Replace **`USER`** with your Linux username:

```
[Unit]
Description="wasmapp node"
After=network-online.target

[Service]
User=USER
ExecStart=/home/USER/go/bin/wasmappd start
Restart=always
RestartSec=3
LimitNOFILE=4096

[Install]
WantedBy=multi-user.target
```

### Start Node Service

Enable and start the service:

```bash
# Enable service
sudo systemctl enable wasmapp.service

# Start service
sudo service wasmapp start

# Check logs
sudo journalctl -fu wasmapp
```

## **Development and Network Connection**

Connect your development environment using the open DNS links provided:

- **API**: https://api-wasmapp.toschdev.com/
- **RPC**: https://rpc-wasmapp.toschdev.com/
- **GRPC**: https://grpc-wasmapp.toschdev.com/

For network details and genesis file, visit: [Network Details & Genesis](https://github.com/toschdev/wasmapp/tree/main/network)

## **Contributing**

Contributions are welcome! Please see the repository's issues for current tasks or submit your suggestions.