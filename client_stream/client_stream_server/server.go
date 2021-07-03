package main

import (
	"fmt"
	"io"
	"log"
	"net"

	"github.com/vremy/go-grpc-examples/client_stream/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	proto.UnimplementedServicesServer
}

func (*server) SendStreamedMessages(stream proto.Services_SendStreamedMessagesServer) error {
	fmt.Println("SendStreamedMessages function was invoked")

	for {
		/** Listen for each request received from the stream and handle it */
		request, err := stream.Recv()

		if err == io.EOF {
			// End of stream
			return stream.SendAndClose(&proto.MessageResponse{
				Result: "Received all streamed messages.",
			})
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
			return err
		}

		fmt.Println(
			"[*] Received request from \""+request.Message.GetName()+"\"",
			" with message \""+request.Message.GetMessage()+"\"",
		)
	}
}

func main() {
	fmt.Println("[*] Client stream server listening for RPC calls...")

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
