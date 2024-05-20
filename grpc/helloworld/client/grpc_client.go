package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "ratwater.xyz/grpc/helloworld/proto"
)

func main() {
	port := flag.Int("port", 5051, "grpc server port")
	greetingName := flag.String("greetingName", "jeff", "greeting name")
	flag.Parse()
	grpcHost := fmt.Sprintf("localhost:%d", *port)

	// Set up a connection to the server.
	conn, err := grpc.NewClient(grpcHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewRatwater_GRPCClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := client.Hello(ctx, &pb.HelloRequest{Name: *greetingName})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetSResponse())

}
