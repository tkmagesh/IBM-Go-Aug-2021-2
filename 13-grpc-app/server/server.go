package main

import (
	"context"
	"fmt"
	"grpc-app/proto"
	"io"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

type server struct {
	proto.UnimplementedAppServiceServer
}

func (s *server) Add(ctx context.Context, req *proto.AddRequest) (*proto.AddResponse, error) {
	x := req.GetX()
	y := req.GetY()
	fmt.Printf("Add request received for X = %d & Y = %d\n", x, y)
	result := x + y
	response := &proto.AddResponse{
		Result: result,
	}
	return response, nil
}

func (s *server) Average(stream proto.AppService_AverageServer) error {
	var count int64
	var sum int64
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			//prepare the response and return the response
			average := sum / count
			response := &proto.AverageResponse{
				Result: average,
			}
			stream.SendAndClose(response)
			return nil
		}
		if err != nil {
			return err
		}
		no := req.GetNo()
		fmt.Printf("Average request received, no = %d\n", no)
		sum += no
		count++
	}
}

func (s *server) GeneratePrime(req *proto.PrimeRequest, stream proto.AppService_GeneratePrimeServer) error {
	fmt.Println("GeneratePrime request received")
	rangeStart := req.GetRangeStart()
	rangeEnd := req.GetRangeEnd()
	for i := rangeStart; i <= rangeEnd; i++ {
		if isPrime(i) {
			time.Sleep(500 * time.Millisecond)
			fmt.Println("Sending Prime No : ", i)
			response := &proto.PrimeResponse{
				No: i,
			}
			stream.Send(response)
		}
	}
	return nil
}

func isPrime(no int64) bool {
	var i int64
	for i = 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}

func (s *server) Greet(stream proto.AppService_GreetServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		greeting := req.GetGreeting()
		first_name := greeting.GetFirstName()
		last_name := greeting.GetLastName()
		greetMsg := fmt.Sprintf("Hi %s %s, Have a nice day!", first_name, last_name)
		time.Sleep(500 * time.Millisecond)
		fmt.Println("Sending : ", greetMsg)
		resp := &proto.GreetResponse{GreetMessage: greetMsg}
		stream.Send(resp)
	}
	return nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln(err)
	}
	grpcServer := grpc.NewServer()
	proto.RegisterAppServiceServer(grpcServer, &server{})
	e := grpcServer.Serve(listener)
	if e != nil {
		log.Fatalln(e)
	}
}
