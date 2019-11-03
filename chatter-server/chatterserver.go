package main

import (
	"strconv"
	chatprotos "chatterbox/chatter-protos"
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
			return nil
		}
		if err != nil {
			log.Fatalf("Error sending client stream %v \n", err)
			return err
		}

		fmt.Printf("Received Message: \n%s\n", mess.Request)

		for i := 0; i < 10; i++ {
			message = "message from server: sending the " + strconv.Itoa(i) + " time\n"

			finalMess := &chatprotos.ChatterBack{
				Response: message,
			}
			stream.Send(finalMess)

			if err != nil {
				return err
			}
		}
	}
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
