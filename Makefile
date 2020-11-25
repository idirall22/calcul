# include .env

gen:
	rm -rf api/pb/
	mkdir api/pb/

	protoc \
	--proto_path=api/proto/ \
	--proto_path=${GOPATH}/src \
	--go_out=api/pb/ \
	--go-grpc_out=api/pb/ api/proto/*.proto

up:
	docker run -d -p 5432:5432 -e POSTGRES_PASSWORD=password --name postgres13 postgres:13-alpine 

down:
	docker rm -f postgres13

run:
	go run cmd/main.go

build:
	docker build -t idirall22/calcul:v1 .








client:
	go run cmd/client/main.go 
	
createdb:
	docker exec -it postgres13 sh -c "createdb -U postgres authz"

dropdb:
	docker exec -it postgres13 sh -c "dropdb -U postgres authz"

connect:
	docker exec -it postgres13 sh -c "psql -U postgres -d authz"

migrate:
	migrate -path db/migrations -database "postgresql://postgres:password@localhost:5432/authz?sslmode=disable" -verbose up

sqlc:
	sqlc generate

store_test: dropdb createdb migrate
	# go test -v pkg/store/postgres/postgres_test.go
	cd pkg/store/postgres && ginkgo

service_test: dropdb createdb migrate
	cd pkg/auth/ && ginkgo
	# go test -v pkg/auth/service_test.go

server:
	go run cmd/server/main.go

.PHONY: up down createdb dropdb connect migrate sqlc test