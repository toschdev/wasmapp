## **Install Go and Cosmovisor**

Feel free to skip this step if you already have Go and Cosmovisor.

### **Install Go**

We will use Go **`v1.21.6`** as example here. The code below also cleanly removes any previous Go installation.

```
sudo rm -rvf /usr/local/go/
wget https://golang.org/dl/go1.21.6.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.21.6.linux-amd64.tar.gz
rm go1.21.6.linux-amd64.tar.gz
```

### **Configure Go**

Unless you want to configure in a non-standard way, then set these in the **`~/.profile`** file.

```
export GOROOT=/usr/local/go
export GOPATH=$HOME/go
export GO111MODULE=on
export PATH=$PATH:/usr/local/go/bin:$HOME/go/bin
```

## **Install Node**

Install the current version of node binary.

```
git clone https://github.com/toschdev/wasmapp wasmapp
cd wasmapp
git checkout v0.1.0
make install
```

## **Configure Node**

### **Initialize Node**

Please replace **`YOUR_MONIKER`** with your own moniker.

```
wasmappd init YOUR_MONIKER --chain-id wasmapp
```

### **Download Genesis**

```
wget -O genesis.json https://raw.githubusercontent.com/toschdev/wasmapp/main/network/genesis.json
mv genesis.json ~/.wasmapp/config
```

### **Configure Seed**

Using a seed node to bootstrap is the best practice.

```
sed -i 's/seeds = ""/seeds = "ade4d8bc8cbe014af6ebdf3cb7b1e9ad36f412c0@seeds.polkachu.com:15156"/' ~/.axelar/config/config.toml

```

## **Launch Node**

### **Create Service File**

Create a **`wasmapp.service`** file in the **`/etc/systemd/system`** folder with the following code snippet. Make sure to replace **`USER`** with your Linux user name. You need sudo previlege to do this step.

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

### **Start Node Service**

```
# Enable service
sudo systemctl enable wasmapp.service

# Start service
sudo service wasmapp start

# Check logs
sudo journalctl -fu wasmapp
```