proto:
	protoc ./pkg/pb/auth.proto --go_out=.
	protoc ./pkg/pb/auth.proto --go-grpc_out=.

server:
	go run cmd/main.go