package main	

import(
	"fmt"
	"log"
	"google.golang.org/grpc"
	"context"
	chatprotos "chatterbox/chatter-protos"
)

func main(){
	fmt.Println("Starting client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
			log.Fatalf("Error %v", err)
	}

	defer cc.Close()

	c := chatprotos.NewChatterboxClient(cc)

	req := "Hello Server, Im the client"

	res , err := c.Chatter(context.Background(), &chatprotos.ChatterThere{
		Request : req,
	})

	fmt.Printf("Client received: %v \n", res.Response)
}