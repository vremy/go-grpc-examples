package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/vremy/go-grpc-examples/unary/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	proto.UnimplementedServicesServer
}

/**
 * Handle the StoreMessage service
 */
func (c *server) StoreMessage(ctx context.Context, request *proto.MessageRequest) (*proto.MessageResponse, error) {
	fmt.Printf("HelloWorld service was invoked with %v\n", request)

	/** Retrieve the name and message from the request */
	name := request.Message.GetName()
	message := request.Message.GetMessage()

	/** Create a response as described by the proto file */
	response := &proto.MessageResponse{
		Result: "Received message: '" + message + "' from '" + name + "'.",
	}

	/** respond back to the client with the response created above */
	return response, nil
}

func main() {
	fmt.Println("[*] Unary server listening for RPC calls...")

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
