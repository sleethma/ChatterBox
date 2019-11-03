package main

import (
	chatprotos "chatterbox/chatter-protos"
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Starting client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error %v", err)
	}

	defer cc.Close()

	c := chatprotos.NewChatterboxClient(cc)

	streamMessage(c)
	
	// fmt.Printf("Client received: %v \n", res.Response)
}

func streamMessage(c chatprotos.ChatterboxClient){
	req := "Hello Server, Im the client"
	stream, err := c.ChatterClientStream(context.Background())
	for i := 0; i < 10; i++ {
		stream.Send(&chatprotos.ChatterThere{
			Request: req,
		})
		time.Sleep(1 * time.Second)
	}
	resp, err2 := stream.CloseAndRecv()
	if err2 != nil {
		log.Printf("Failed to close and send from client: %v \n", err)
	}
	fmt.Printf("Received:\n%v ", resp.Response)
}
