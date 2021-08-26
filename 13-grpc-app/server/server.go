package main

import (
	"context"
	"fmt"
	"grpc-app/proto"
	"io"
	"log"
	"net"

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
