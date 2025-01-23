# GoPong

## Description
GoPong is a multiplayer version of the classic Pong game, implemented in the terminal using the Go programming language and gRPC technology.

The game allows multiple players to connect to a server and compete in real-time. The server handles the game logic, synchronizing player actions, while the client provides a minimalist interface for controlling and displaying the game directly in the terminal.

---

## Environment configuration
Before running the application, configure your environment variables.

Use the provided [.env.example](./.env.example) file as a reference.

---

## Requirements for Building and Running Locally
Ensure you have the following installed on your system:
1. [Go](https://go.dev/dl/) – for building the application.
2. [task](https://github.com/go-task/task) – for automating build and run commands.
    ```sh
    go install github.com/go-task/task/v3/cmd/task@latest
    ```

---

## Running the Client Locally
The client must be built and run on your local machine. Follow these steps:

Build all .proto and the Client:
```sh
task generate
task build-client
```

Run the Client:
```sh
task run-client
```

---

## Running the Server with Docker
The server can be built and run using Docker. Follow these steps:

TODO: Create Docker configuration for the Server.

---

## Running the Server Locally
Build all .proto and the Server:
```sh
task generate
task build-server
```

Run the Server:
```sh
task run-server
```

Alternatively you can generate .proto and build whole project(Client + Server):
```sh
task build
```

---

## Frameworks
- [gRPC](https://github.com/grpc/grpc-go) – enables efficient client-server communication.
- [protobuf](https://github.com/protocolbuffers/protobuf) – used for data serialization to ensure lightweight and fast communication.
- [Bubbletea](https://github.com/charmbracelet/bubbletea) – provides a modern terminal-based UI framework for the client.
