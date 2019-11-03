package main

import (
	chatprotos "chatterbox/chatter-protos"
	"context"
	"fmt"
	"io"
	"log"
	"strconv"

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

	go sendStream(stream)
	go receiveStream(stream, waitCh)

	<-waitCh
}

func sendStream(stream chatprotos.Chatterbox_ChatterClientStreamClient) error {

	for i := 0; i < 10; i++ {
		message := "Hello Server, I am " + strconv.Itoa(i) + " client\n "

		stream.Send(&chatprotos.ChatterThere{
			Request: message,
		})

	}
	stream.CloseSend()
	return nil
}

func receiveStream(stream chatprotos.Chatterbox_ChatterClientStreamClient, wait chan struct{}) {

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			close(wait)
		}
		if err != nil {
			fmt.Printf("Error establishing stream %v\n", err)
		}
		fmt.Printf("Client recieved this response: %v", resp.Response)

	}
}
