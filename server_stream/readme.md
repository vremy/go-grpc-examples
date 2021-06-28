# Server stream GRPC service example
This example demonstrate a GRPC server and client communicating with each other
in server stream mode. This means that a client send a single request to the
server and the server streams back multiple responses.

In this example a simple message is send containing the name of the sender and
the message itself. The server will respond by telling the client it received
the message from the sender in 5 seperate response through streaming.

## Usage
1. Open two terminals in split screen mode in the unary directory;
1. Run the server in the left screen by running the following command.
    ```sh
    go run server_stream_server/server.go
    ```
1. Run the client in the right screen by running the following command;
    ```sh
    go run server_stream_server/server.go
    ```

