# go-simple-bank
under development

add mockgen to path

sudo nano ~/.bashrc
```
export PATH=$PATH:~/go/bin
```

Create new db migration:
```bash
migrate create -ext sql -dir db/migration -seq <migration_name>
```

protoc
evans
swagger
statik


Start Server
make startdb
make serve


Create new query
make sqlc
make mock

Create gRPC entrypoint
create new rpc proto file
add rpc proto file import to service_simple_bank.proto
generate protoc code with make proto
create gapi go file
