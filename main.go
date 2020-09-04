package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	hello "grpcDemo/hello/demo"
	"log"
	"net"
	"strconv"
	"time"
)

type DemoService struct {

}

func (m *DemoService) Hello(ctx context.Context,req *hello.String) (*hello.String, error) {
	fmt.Println("hello")
	fmt.Println(req.Name)
	return &hello.String{Name:strconv.Itoa(int(time.Now().Unix())),Age:11}, nil
}


func main(){
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50001))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	hello.RegisterSayHelloServiceServer(grpcServer, &DemoService{})
	reflection.Register(grpcServer)
	grpcServer.Serve(lis)
}

//
//func interceptor(ctx context.Context,req interface{},info *grpc.UnaryServerInfo, handler grpc.UnaryHandler)(resp interface{},err error){
//	fmt.Println("log")
//	fmt.Println(info.FullMethod)
//	return handler(ctx,req)
//}