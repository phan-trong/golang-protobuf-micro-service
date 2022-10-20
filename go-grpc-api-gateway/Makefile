proto:
	protoc ./pkg/auth/pb/auth.proto --go_out=.
	protoc ./pkg/auth/pb/auth.proto --go-grpc_out=.
	protoc ./pkg/order/pb/order.proto --go_out=.
	protoc ./pkg/order/pb/order.proto --go-grpc_out=.
	protoc ./pkg/product/pb/product.proto --go_out=.
	protoc ./pkg/product/pb/product.proto --go-grpc_out=.
server:
	go run cmd/main.go