DB_URL=postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable

network:
	docker network create bank-network
	
postgres:
	docker run --name postgres14 --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14-alpine

startdb:
	docker start postgres14

createdb:
	docker exec -it postgres14 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres14 dropdb simple_bank

migrateup:
	migrate -path db/migration/ -database "${DB_URL}" -verbose up

migrateup1:
	migrate -path db/migration/ -database "${DB_URL}" -verbose up 1

migratedown:
	migrate -path db/migration/ -database "${DB_URL}" -verbose down

migratedown1:
	migrate -path db/migration/ -database "${DB_URL}" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go  github.com/xdot2012/simple-bank/db/sqlc Store 

proto:
	rm -f pb/*.go
	rm -f doc/swagger/*.swagger.json
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt paths=source_relative \
	--openapiv2_out=doc/swagger --openapiv2_opt=allow_merge=true,merge_file_name=simple_bank \
    proto/*.proto
	statik -src=./doc/swagger -dest=./doc

evans:
	evans --host localhost --port 8001 -r repl

redis:
	docker run --name redis -p 6379:6379 -d redis:7-alpine

startredis:
	docker start redis

redis-ping:
	docker exec -it redis redis-cli ping

.PHONY: print postgres createdb dropdb migrateup migrateup1 migratedown migratedown1 sqlc test server mock proto redis