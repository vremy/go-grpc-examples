# Client stream GRPC service example
This example demonstrate a GRPC server and client communicating with each other
in client stream mode. This means that a client streams multiple request to the
server. The server responds one time to indicate to the client it received all
requests.

In this example multiple message request are streamed from the client to the
server. Each streamed message request contains a name and the message itself.

## Usage
1. Open two terminals in split screen mode in the client_stream directory;
1. Run the server in the left screen by running the following command.
    ```sh
    go run client_stream_server/srver.go 
    ```
1. Run the client in the right screen by running the following command;
    ```sh
    go run client_stream_client/client.go
    ```

