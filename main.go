package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/shishir127/golang-grpc-client/spike"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main() {
	port := os.Getenv("PORT")
	accessToken := os.Getenv("TOKEN")
	serverAddr := flag.String("server_addr", "127.0.0.1:"+port, "The server address in the format of host:port")

	conn, err := grpc.Dial(*serverAddr, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Error while establishing connection")
		fmt.Println(err)
		return
	}

	defer conn.Close()

	client := spike.NewStreamerClient(conn)

	request := &spike.HelloRequest{Name: "Shishir"}
	md := metadata.Pairs("Authorization", accessToken)
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	stream, err := client.SayHello(ctx, request)
	if err != nil {
		fmt.Println("Error while streaming")
		fmt.Println(err)
		return
	}

	fmt.Println("Reading stream from server")
	for {
		reply, err := stream.Recv()
		if err == io.EOF {
			break
		}
		fmt.Println(reply)
		if err != nil {
			fmt.Println(reply)
		} else {
			fmt.Println("Error in receiving message in stream")
			fmt.Println(err)
		}
	}
}
