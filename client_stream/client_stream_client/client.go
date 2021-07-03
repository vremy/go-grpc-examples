package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/vremy/go-grpc-examples/client_stream/proto"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("[*] Starting to do a Client Stream RPC...")

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

	requestClientStreamService(client)
}

func requestClientStreamService(client proto.ServicesClient) {
	/** Give the server 5 seconds to respond or else cancel the request. */
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	/** Open a stream to the SendStreamMessages service */
	stream, err := client.SendStreamedMessages(ctx)

	if err != nil {
		log.Fatalf(
			"Error occurred after calling SendStreamedMessages service: %v",
			err,
		)
	}

	/** Loop through each request and send it through the opened stream */
	for _, request := range getRequests() {
		fmt.Printf("Sending request: %v\n", request)
		stream.Send(request)
		time.Sleep(1 * time.Second)
	}

	/** Close stream and receive response */
	response, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf(
			"Error while receiving response from SendStreamedMessages: %v",
			err,
		)
	}

	fmt.Printf(
		"Received response from SendStreamedMessages service: %v\n",
		response,
	)
}

func getRequests() []*proto.MessageRequest {
	return []*proto.MessageRequest{
		{
			Message: &proto.Message{
				Name:    "Foo",
				Message: "This is the first message.",
			},
		},
		{
			Message: &proto.Message{
				Name:    "Foo Bar",
				Message: "This is the second message.",
			},
		},
		{
			Message: &proto.Message{
				Name:    "Bar",
				Message: "This is the thirth message.",
			},
		},
		{
			Message: &proto.Message{
				Name:    "Bar Foo",
				Message: "This is the fourth message.",
			},
		},
	}
}
