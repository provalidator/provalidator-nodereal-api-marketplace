# Install watcher
```
go install github.com/cosmtrek/air@latest
air init
air -c .air.toml
```

# Air
```
# Swagger & main build
cmd = "swag init --instanceName=lcd --generalInfo=routes/lcd.go --exclude=routes/rpc --output docs/lcd && swag init --instanceName=rpc --generalInfo=routes/rpc.go --exclude=routes/lcd --output docs/rpc && go build -o ./tmp/main.exe ."
or 
# Main build only
cmd = "go build -o ./tmp/main.exe ."
```

# Service
```
echo "[Unit]
Description=rpc daemon
After=network-online.target

[Service]
User=$USER
WorkingDirectory=$HOME/go/src/github.com/provalidator-nodereal-api-marketplace
ExecStart=$HOME/go/src/github.com/provalidator-nodereal-api-marketplace/main

Restart=always
RestartSec=3
LimitNOFILE=65535

[Install]
WantedBy=multi-user.target
" >provalidator-nodereal-api-marketplace.service

sudo mv provalidator-nodereal-api-marketplace.service /etc/systemd/system/provalidator-nodereal-api-marketplace.service
sudo -S systemctl daemon-reload
sudo systemctl enable provalidator-nodereal-api-marketplace
sudo systemctl start provalidator-nodereal-api-marketplace
journalctl -fu provalidator-nodereal-api-marketplace
sudo systemctl stop provalidator-nodereal-api-marketplace
```