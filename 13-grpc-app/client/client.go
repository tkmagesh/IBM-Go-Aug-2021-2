package main

import (
	"context"
	"fmt"
	"grpc-app/proto"
	"log"
	"time"

	"google.golang.org/grpc"
)

func main() {
	client, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()
	clientConn := proto.NewAppServiceClient(client)
	//doRequestResponse(clientConn)
	doClientStreaming(clientConn)
}

func doRequestResponse(clientConn proto.AppServiceClient) {
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

func doClientStreaming(clientConn proto.AppServiceClient) {
	var nos []int64 = []int64{8, 3, 5, 2, 6, 4, 1, 7, 9}
	clientStream, err := clientConn.Average(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	for _, no := range nos {
		time.Sleep(time.Millisecond * 500)
		fmt.Println("Sending : ", no)
		req := &proto.AverageRequest{
			No: no,
		}
		clientStream.Send(req)
	}
	res, e := clientStream.CloseAndRecv()
	if e != nil {
		log.Fatalln(e)
	}
	fmt.Println("Received : ", res.GetResult())
}
