package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/vremy/go-grpc-examples/unary/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	fmt.Println("[*] Starting to do a Unary RPC...")

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

	requestUnaryService(client)
}

func requestUnaryService(client proto.ServicesClient) {
	/**
	 * Create the request as described in the proto file. This request contains
	 * a name and message as shown below.
	 */
	request := &proto.MessageRequest{
		Message: &proto.Message{
			Name:    "Foo",
			Message: "Hello unary service!",
		},
	}

	/** Give the server 5 seconds to respond or else cancel the request. */
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	/**
	 * Send the request to service StoreMessage and receive a response or
	 * error from the server.
	 */
	response, err := client.StoreMessage(ctx, request)

	if err != nil {
		statusError, ok := status.FromError(err)
		/**
		 * Check if the error is thrown because of a request timeout. When this
		 * is the case it will show an message about the timeout, else it will
		 * just display the error itself.
		 */
		if ok && statusError.Code() == codes.DeadlineExceeded {
			log.Println("[!] Request timeout occurred.")
			return
		}

		log.Fatalf("[!] Error occurred while calling RPC service: %v", err)
	}

	/** if everything whent well we just output the response from the server */
	log.Printf("[*] Response from StoreMessage: %v", response.Result)
}
