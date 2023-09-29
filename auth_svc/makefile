auth:
	protoc --go_out=. --go-grpc_out=. --proto_path=./pkg/auth/pb/ auth.proto

product:
	protoc --go_out=. --go-grpc_out=. --proto_path=./pkg/product/pb/ product.proto

run:
	go run ./cmd/api/main.go

