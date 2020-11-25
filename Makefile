gen:
	rm -rf api/pb/
	mkdir api/pb/

	protoc \
	--proto_path=api/proto/ \
	--proto_path=${GOPATH}/src \
	--go_out=api/pb/ \
	--go-grpc_out=api/pb/ api/proto/*.proto

run:
	go run cmd/main.go

client:
	go run cmd/client/main.go 

build:
	docker build -t idirall22/calcul:v1 .

publish:
	docker push idirall22/calcul:v1


.PHONY: run client build publish