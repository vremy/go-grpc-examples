package main

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"time"

	"github.com/vremy/go-grpc-examples/server_stream/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	proto.UnimplementedServicesServer
}

/**
 * Handle the GetStreamedMessages service
 */
func (c *server) GetStreamedMessages(request *proto.MessageRequest, stream proto.Services_GetStreamedMessagesServer) error {
	/** Retrieve the name and message from the request */
	name := request.Message.GetName()

	fmt.Printf("[!] GetStreamedMessages service was invoked with %v\n", request)

	for i := 0; i < 5; i++ {
		/** Create a response as described by the proto file */
		response := &proto.MessageResponse{
			Result: "Hello " + name + " this is response " + strconv.Itoa(i),
		}

		/** Stream each response back to the client with an interval of 1 sec */
		stream.Send(response)
		time.Sleep(1 * time.Second)
	}

	return nil
}

func main() {
	fmt.Println("[*] Stream server listening for RPC calls...")

	/**
	 * Create a TCP listener listening on 0.0.0.0:50051. Port 50051 is the
	 * default port for GRPC services.
	 */
	listener, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	/** Create a new GRPC server */
	grpcServer := grpc.NewServer()

	/** Enable reflection on the GRPC server */
	reflection.Register(grpcServer)

	/** Make the services described in the proto file available for the server */
	proto.RegisterServicesServer(grpcServer, &server{})

	/** Trow an error when listening whent wrong */
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
