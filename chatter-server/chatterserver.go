package main

import (
	chatprotos "chatterbox/chatter-protos"
	"context"
	"fmt"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
}

func (*server) ChatterClientStream(stream chatprotos.Chatterbox_ChatterClientStreamServer) error {
	message := ""
	for {
		mess, err := stream.Recv()
		if err == io.EOF {
			finalMess := &chatprotos.ChatterBack{
				Response: message,
			}
			fmt.Printf("Received Message: %s\n", finalMess.Response)
			stream.SendAndClose(finalMess)
			break
		}
		if err != nil {
			log.Fatalf("Error sending client stream %v \n", err)
			return err
		}
		message += "New message: " + mess.Request + "\n"
	}
	return nil
}

func (*server) Chatter(ctx context.Context, req *chatprotos.ChatterThere) (*chatprotos.ChatterBack, error) {
	message := "Heard you say: " + req.Request

	res := &chatprotos.ChatterBack{
		Response: message,
	}
	fmt.Printf("Received response from client %v\n", res)

	return res, nil
}

func main() {
	fmt.Println("Starting server")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	cc := grpc.NewServer()

	chatprotos.RegisterChatterboxServer(cc, &server{})

	err2 := cc.Serve(lis)
	if err != nil {
		log.Fatalf("Err serving from listener %v\n", err2)
	}

}
