package main

import (
	"log"
	"google.golang.org/grpc"
	chatprotos "chatterbox/chatter-protos"
"net"
"fmt"
"context"
)


type servers struct{

}

func (*servers) Chatback(ctx context.Context, req *chatprotos.ChatterThere) (*chatprotos.ChatterBack, error){

	message := "Heard you say: " + req.Request 
	
	res := &chatprotos.ChatterBack{
		Response : "Heard",
	}
	return nil, nil
}

func main(){
	fmt.Println("Starting server")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	cc := grpc.NewServer()

	chatprotos.RegisterChatterboxServer(cc, &servers)

	err2 := cc.Serve(lis)
	if err != nil{
		log.Fatalf("Err serving from listener %v\n", err2)
	}

}