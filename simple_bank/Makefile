postgres:
	docker run --name  postgres15 -p 5432:5432 --network bank-network -e POSTGRES_USER=root -e POSTGRES_PASSWORD=assim -d postgres:15.3-alpine3.18

createdb:
	docker exec -it postgres15 createdb --username=root --owner=root simple_bank

dropdb:
	winpty docker exec -it postgres14 dropdb simple_bank

migrateup:
	migrate -path db/migrations  -database "postgresql://root:assim@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://root:assim@localhost:5432/simple_bank?sslmode=disable" -verbose down

migrateup1:
	migrate -path db/migrations  -database "postgresql://root:assim@localhost:5432/simple_bank?sslmode=disable" -verbose up 1

migratedown1:
	migrate -path db/migrations -database "postgresql://root:assim@localhost:5432/simple_bank?sslmode=disable" -verbose down 1

db_docs:
	dbdocs build doc/db.dbml

db_schema:
	dbml2sql --postgres -o doc/schema.sql doc/db.dbml
sqlc: 
	docker run --rm -v "%cd%:/src" -w /src kjconroy/sqlc generate

test:
	go test -v -cover -short ./...

server:
	go run main.go

mock:
	mockgen -package mockdb  -destination db/mock/store.go github.com/prepStation/simple_bank/db/sqlc Store
	mockgen -package mockwk  -destination worker/mock/distributor.go github.com/prepStation/simple_bank/worker TaskDistributor

proto:
	rm -f pb/*.go 
	rm -f doc/swagger/*.swagger.json 
	rm -f doc/statik/*.go
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
	--openapiv2_out=doc/swagger --openapiv2_opt=allow_merge=true,merge_file_name=simple_bank  \
    proto/*.proto
	statik -src=./doc/swagger -dest=./doc

evans:
	evans --host localhost --port 9090 -r repl

redis:
	docker run --name redis -p 6379:6379 -d redis:7-alpine

run_migration:
	migrate create -ext sql -dir db/migrations -seq $(name)

.PHONY: postgres createdb dropdb redis migrateup evans migratedown sqlc run_migration server proto mock migrateup1 migratedown1 db_docs db_schema