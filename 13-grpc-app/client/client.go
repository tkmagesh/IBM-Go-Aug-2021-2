package main

import (
	"context"
	"fmt"
	"grpc-app/proto"
	"log"

	"google.golang.org/grpc"
)

func main() {
	client, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()
	clientConn := proto.NewAppServiceClient(client)
	req := &proto.AddRequest{
		X: 100,
		Y: 200,
	}
	resp, e := clientConn.Add(context.Background(), req)
	if e != nil {
		log.Fatalln(e)
	}
	fmt.Println(resp.GetResult())
}
