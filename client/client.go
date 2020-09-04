package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"grpcDemo/client/resolver"
	"grpcDemo/hello/demo"
	"strconv"
	"time"
)
func main(){
	conn,err:=grpc.Dial("demo:///demo-svc",
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithBalancerName(roundrobin.Name),
		grpc.WithResolvers(resolver.NewResolverBuilder()),
		//grpc.WithKeepaliveParams(
		//	keepalive.ClientParameters{
		//		Time:time.Second*18,
		//		Timeout:time.Second*1,
		//		PermitWithoutStream:true,
		//	}),
		grpc.WithUnaryInterceptor(
		UnaryClientInterceptor))
	if err!=nil{
		panic(err)
	}
	_,err=grpc.Dial("demo:///demo-svc",
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithBalancerName(roundrobin.Name),
		grpc.WithResolvers(resolver.NewResolverBuilder()),
		//grpc.WithKeepaliveParams(
		//	keepalive.ClientParameters{
		//		Time:time.Second*18,
		//		Timeout:time.Second*1,
		//		PermitWithoutStream:true,
		//	}),
		grpc.WithUnaryInterceptor(
			UnaryClientInterceptor))
	if err!=nil{
		panic(err)
	}
	//client.Hello(context.TODO(),&hello.String{})
	client:= hello.NewSayHelloServiceClient(conn)
	client= hello.NewSayHelloServiceClient(conn)
	for {
		client.Hello(context.TODO(),&hello.String{Name:strconv.Itoa(int(time.Now().Unix()))})
		time.Sleep(time.Second*3)
	}

	time.Sleep(time.Second*100)

}
func UnaryClientInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	//log.Printf("before invoker. method: %+v, request:%+v", method, req)
	err := invoker(ctx, method, req, reply, cc, opts...)
	//log.Printf("after invoker. reply: %+v", reply)
	return err
}