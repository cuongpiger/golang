protoc:
	protoc \
		--go_out=Mproto/upload.proto=.:proto \
		--go-grpc_out=Mproto/upload.proto=.:proto \
		proto/upload.proto

runServer:
	go run cmd/server/main.go
