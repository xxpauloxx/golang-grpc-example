# Golang GRPC example

Example implementation in Golang of GRPC, below is the description of how to build and run it.

Build the GRPC server:

```sh
$ cd server
$ protoc --go_out=. --go-grpc_out=. helloworld.proto
$ go run server.go
```

Build the GRPC client:

```sh
$ cd server
$ protoc --go_out=. --go-grpc_out=. helloworld.proto
$ go run server.go
```

:)
