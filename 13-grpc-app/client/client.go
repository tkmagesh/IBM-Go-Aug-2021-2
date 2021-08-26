package main

import (
	"context"
	"fmt"
	"grpc-app/proto"
	"io"
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
	//doClientStreaming(clientConn)
	//doServerStreaming(clientConn)
	doBiDirectionalStreaming(clientConn)
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

func doServerStreaming(clientConn proto.AppServiceClient) {
	req := &proto.PrimeRequest{
		RangeStart: 5,
		RangeEnd:   100,
	}
	stream, err := clientConn.GeneratePrime(context.Background(), req)
	if err != nil {
		log.Fatalln(err)
	}
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("Received : ", resp.GetNo())
	}
}

func doBiDirectionalStreaming(client proto.AppServiceClient) {
	stream, err := client.Greet(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	requestData := []*proto.GreetRequest{
		&proto.GreetRequest{
			Greeting: &proto.Greeting{
				FirstName: "Magesh",
				LastName:  "Kuppan",
			},
		},
		&proto.GreetRequest{
			Greeting: &proto.Greeting{
				FirstName: "Suresh",
				LastName:  "Kannan",
			},
		},
		&proto.GreetRequest{
			Greeting: &proto.Greeting{
				FirstName: "Ramesh",
				LastName:  "Jayaraman",
			},
		},
		&proto.GreetRequest{
			Greeting: &proto.Greeting{
				FirstName: "Rajesh",
				LastName:  "Pandit",
			},
		},
		&proto.GreetRequest{
			Greeting: &proto.Greeting{
				FirstName: "Naresh",
				LastName:  "Kumar",
			},
		},
	}

	go func() {
		for _, req := range requestData {
			stream.Send(req)
		}
		stream.CloseSend()
	}()
	/* wg := &sync.WaitGroup{}
	wg.Add(1) */
	done := make(chan bool)
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				fmt.Println("Thats all folks!")
				break
			}
			if err != nil {
				log.Fatalln(err)
			}
			log.Println("Greet Result:", res.GetGreetMessage())
		}
		//done <- true
		close(done)
	}()
	<-done
}
