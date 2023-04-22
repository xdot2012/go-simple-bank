# go-simple-bank

## Links
https://protobuf.dev/getting-started/gotutorial/

## Installation

### GO
Install go (https://go.dev/dl/)

#### EDIT BASHRC
sudo nano ~/.bashrc
```
export GO_PATH=$PATH:~/go
export PATH=$PATH:/$GO_PATH/bin
```

### DOCKER
Install docker (https://docs.docker.com/engine/install/ubuntu/)

### SQLC
sudo snap install sqlc

### GOLANG-MIGRATE
Download and install golang-migrate(https://github.com/golang-migrate/migrate/releases)

### PROTOC
Download and install protoc(https://github.com/protocolbuffers/protobuf/releases)<br />
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest<br />
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest<br />
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway<br />
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2<br />
github.com/rakyll/statik<br />

### EVANS
Download and install evans(https://github.com/ktr0731/evans/releases)

### MOCKGEN
go install github.com/golang/mock/mockgen@v1.6.0

### RUN MAKEFILE COMMANDS
make network<br />
make postgres<br />
make createdb<br />
make redis<br />

# Development
## Run Server:
make startdb<br />
make startredis<br />
make server<br />


## Create new db migration:
```bash
migrate create -ext sql -dir db/migration -seq <migration_name>
```


## Tools
protoc<br />
evans<br />
swagger<br />
statik<br />
dbdiagrams.io<br />


## How to create new database query:
1. create query on db/query/
2. run commands:
```
make sqlc
make mock
```

## How to create new gRPC entrypoint:
Create gRPC entrypoint
1. create new rpc proto file
2. add rpc proto file import to service_simple_bank.proto
3. generate protoc code with make proto
3. create gapi go file


## Add New Environment Variable
1. change app.env
2. update util/config.go

## Log tools
Logstash
Fluentd
Grafana loki