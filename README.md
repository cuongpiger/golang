###### [_↩ Back to `main` branch_](https://github.com/cuongpiger/golang)

<hr>

# gRPC with Buf and Go

- Install `grpcurl`:

```bash
brew install grpcurl
```

- Start the gRPC server:

```bash
make runServer
```

- List all services:

```bash
grpc_cli ls localhost:50051
```

- Make the request to `EchoService`:

```bash
grpc_cli call localhost:50051 Echo "message: 'Hello, Cuong Duong'"
```

![](./assets/01.png)
