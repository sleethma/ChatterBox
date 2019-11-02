package main

import (
	"log"
	"google.golang.org/grpc"
	chatprotos "chatterbox/chatter-protos"
"net"
"fmt"
"context"
)


type server struct{

}

func (*server) Chatter(ctx context.Context, req *chatprotos.ChatterThere) (*chatprotos.ChatterBack, error){
	message := "Heard you say: " + req.Request 
	
	res := &chatprotos.ChatterBack{
		Response : message,
	}
	fmt.Printf("Received response from client %v\n", res)

	return res, nil
}

func main(){
	fmt.Println("Starting server")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	cc := grpc.NewServer()

	chatprotos.RegisterChatterboxServer(cc, &server{})

	err2 := cc.Serve(lis)
	if err != nil{
		log.Fatalf("Err serving from listener %v\n", err2)
	}

}