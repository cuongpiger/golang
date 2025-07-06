buf:
	buf generate

runServer:
	@go run main.go

.PHONY: buf runServer