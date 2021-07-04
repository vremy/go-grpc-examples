# Bi-Directional stream GRPC service example
This example demonstrate a GRPC server and client communicating with each other
in bi-directional stream mode. This means that a client streams multiple
requests to the server and the server streams multiple responses back to the
client.

## Usage
1. Open two terminals in split screen mode in the server_stream directory;
1. Run the server in the left screen by running the following command.
    ```sh
    go run bi_directional_stream/bi_directional_stream_server/server.go
    ```
1. Run the client in the right screen by running the following command;
    ```sh
    go run bi_directional_stream/bi_directional_stream_client/client.go
    ```

