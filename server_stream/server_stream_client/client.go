package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/vremy/go-grpc-examples/server_stream/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	fmt.Println("[*] Starting to do a server stream RPC call...")

	/**
	 * Let the GRPC client connect over an INSECURE connection with the GRPC
	 * server over port 50051.
	 */
	grpcClient, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	/** Close the connection at the end of this function */
	defer grpcClient.Close()

	/**
	 * Make the services configured in the proto file accessible by the client
	 */
	client := proto.NewServicesClient(grpcClient)

	requestServerStreamService(client)
}

func requestServerStreamService(client proto.ServicesClient) {
	/**
	 * Create the request as described in the proto file. This request contains
	 * a name and message as shown below.
	 */
	request := &proto.MessageRequest{
		Message: &proto.Message{
			Name:    "Foo",
			Message: "Hello GetStreamedMessages service!",
		},
	}

	/** Give the server 10 seconds to respond or else cancel the request. */
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	/**
	 * Send the request to service GetStreamedMessages and receive a response or
	 * error from the server.
	 */
	responses, err := client.GetStreamedMessages(ctx, request)

	if err != nil {
		log.Fatalf("Error while calling GetStreamedMessages RPC: %v", err)
		return
	}

	for {
		/** Iterate over each received response from the stream */
		response, err := responses.Recv()

		if err == io.EOF {
			break // End of stream
		}

		if err != nil {
			statusError, ok := status.FromError(err)
			/**
			 * Check if the error is thrown because of a request timeout. When
			 * this is the case it will show an message about the timeout, else
			 * it will just display the error itself.
			 */
			if ok && statusError.Code() == codes.DeadlineExceeded {
				log.Println("[!] Request timeout occurred.")
				return
			}

			log.Fatalf("[!] Error occurred while calling RPC service: %v", err)
		}

		/** Display each received response from the stream */
		log.Printf(
			"Response from GetStreamedMessages RPC service: %v",
			response.GetResult(),
		)
	}
}
