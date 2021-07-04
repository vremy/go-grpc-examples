package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/vremy/go-grpc-examples/bi_directional_stream/proto"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("[*] Starting to do a Bi-Directional Stream RPC...")

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

	requestStreamedMessages(client)
}

func requestStreamedMessages(client proto.ServicesClient) {
	/** Give the server 10 seconds to respond or else cancel the request. */
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	/** Open a stream to the StreamedMessages service */
	stream, err := client.StreamedMessages(ctx)

	if err != nil {
		log.Fatalf(
			"Error occurred after calling StreamedMessages service: %v",
			err,
		)
	}

	// Send messages through a stream in a Goroutine
	go streamMessages(stream)

	/**
	 * Setup channel to make the client wait until all responses are processed
	 * by the goroutine.
	 */
	wait := make(chan struct{})
	go receiveStreamedMessages(stream, wait)
	<-wait
}

func streamMessages(stream proto.Services_StreamedMessagesClient) {
	for _, request := range getRequests() {
		fmt.Printf("Sending message: %v\n", request)
		stream.Send(request)
		time.Sleep(1000 * time.Millisecond)
	}
	stream.CloseSend()
}

func receiveStreamedMessages(
	stream proto.Services_StreamedMessagesClient,
	wait chan struct{},
) {
	for {
		/** Read each response from the stream */
		response, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while receiving: %v", err)
			break
		}

		fmt.Printf("Received: %v\n", response.GetResult())
	}

	close(wait)
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
