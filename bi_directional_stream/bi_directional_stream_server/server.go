package main

import (
	"fmt"
	"io"
	"log"
	"net"

	"github.com/vremy/go-grpc-examples/bi_directional_stream/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	proto.UnimplementedServicesServer
}

func (*server) StreamedMessages(stream proto.Services_StreamedMessagesServer) error {
	fmt.Println("StreamedMessages function was invoked.")

	for {
		/** Read each request from the stream */
		request, err := stream.Recv()

		if err == io.EOF {
			// End of stream
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
			return err
		}

		/** After receiving a request send a response back to the client */
		sendErr := stream.Send(&proto.MessageResponse{
			Result: "Hello " + request.Message.GetName() + " ",
		})

		if sendErr != nil {
			log.Fatalf("Error while sending data to client: %v", sendErr)
			return sendErr
		}
	}
}

func main() {
	fmt.Println("[*] Bi-Directional stream server listening for RPC calls...")

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
