# Unary GRPC service example
This example demonstrate a GRPC server and client communicating with each other
in unary mode. This means that a unary client behaves in the same manner as a
REST api will behave. So a single request will be send by the client and a 
single response will be given by the server. For this communication no streaming
will be used.

In this example a simple message is send containing the name of the sender and
the message itself. The server will respond by telling the client it received
the message from the sender.

## Usage
1. Open two terminals in split screen mode in the unary directory;
1. Run the server in the left screen by running the following command.
    ```sh
    go run unary_server/server.go
    ```
1. Run the client in the right screen by running the following command;
    ```sh
    go run unary_client/client.go
    ```

