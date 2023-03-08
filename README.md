# Golang gRPC Example

This is a simple example of how to create a gRPC server and client using Go.

### Requirements
- Go (version 1.14 or higher)
- Protocol Buffers Compiler (protoc)

### Installation

Clone this repository to your local machine using git clone https://github.com/xxpauloxx/golang-grpc-example.git.
Install the Protocol Buffers Compiler by following the instructions for your operating system here.  

Generate the Go code for the protocol buffer messages and gRPC server and client stubs by running the following command from the root directory of the project:

```sh
$ cd <server> or <client>
$ protoc --go_out=. --go-grpc_out=. helloworld.proto
```

### Usage
Start the server by running the following command from the server directory of the project:

```sh
$ go run server.go
```

Run the client by running the following command from the client directory of the project:

```sh
$ go run client.go
```

The client will make a request to the server, which will respond with a simple greeting message.

### Contributing
If you find a bug or have an idea for a new feature, feel free to open an issue or submit a pull request. Contributions are always welcome!

### License
This project is licensed under the MIT License - see the LICENSE file for details.
