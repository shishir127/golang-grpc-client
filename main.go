package main

import (
	"context"
	"flag"
	"fmt"
	"io"

	"github.com/shishir127/golang-grpc-client/spike"
	"google.golang.org/grpc"
)

func main() {
	serverAddr := flag.String("server_addr", "127.0.0.1:50051", "The server address in the format of host:port")

	conn, err := grpc.Dial(*serverAddr, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Error while establishing connection")
		fmt.Println(err)
		return
	}

	defer conn.Close()

	client := spike.NewStreamerClient(conn)

	request := &spike.HelloRequest{Name: "Shishir"}
	stream, err := client.SayHello(context.Background(), request)
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
