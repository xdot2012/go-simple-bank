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

protoc<br />
evans<br />
swagger<br />
statik<br />


Start Server
```
make startdb
make serve
```

Create new call to database
1. create query on db/query/
2. run commands:
```
make sqlc
make mock
```

Create gRPC entrypoint
1. create new rpc proto file
2. add rpc proto file import to service_simple_bank.proto
3. generate protoc code with make proto
3. create gapi go file
