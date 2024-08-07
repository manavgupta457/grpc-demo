package main

import (
	"log"

	pb "github.com/grpc-demo/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)
	names := &pb.NamesList{
		Names: []string{"Akhil", "Alice", "Bob"},
	}

	//The below is for unary rpc
	// callSayHello(client)

	//The below is for server streaming
	// callSayHelloServerStream(client, names)

	//The below is for client streaming
	// callSayHelloClientStreaming(client, names)

	//The below is for bidirectional streaming
	callHelloBidirectionalStream(client, names)

}
