proto:
	protoc ./pkg/pb/product.proto --go_out=.
	protoc ./pkg/pb/product.proto --go-grpc_out=.

server:
	go run cmd/main.go