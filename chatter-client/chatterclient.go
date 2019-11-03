package main

import (
	"io"
	"strconv"
	chatprotos "chatterbox/chatter-protos"
	"context"
	"fmt"
	"log"
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

	biDiStreaming(c)
}

func biDiStreaming(c chatprotos.ChatterboxClient) {

	stream, err := c.ChatterClientStream(context.Background())
	if err != nil {
		fmt.Printf("Error establishing stream %v\n", err)
		return
	}

	waitCh := make(chan struct{})
	
	go func() {
		for i:=0;i<10;i++{
		message := "Hello Server, I am " + strconv.Itoa(i) + " client"
		stream.Send(&chatprotos.ChatterThere{
			Request: message,
		})
	}
	stream.CloseSend()
	}()

	go func() {
		for{
		resp, err := stream.Recv()
		if err == io.EOF{
			close(waitCh)
			return
		}
		if err != nil {
			fmt.Printf("Error establishing stream %v\n", err)
			return
		}
		fmt.Printf("I have recieved this response: %v", resp.Response)
	}
	}()

	<-waitCh
}


