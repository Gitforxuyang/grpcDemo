package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/fullstorydev/grpcurl"
	"github.com/gdong42/grpc-mate/metadata"
	"github.com/gdong42/grpc-mate/proxy/reflection"
	"github.com/gdong42/grpc-mate/proxy/stub"
	"github.com/jhump/protoreflect/dynamic/grpcdynamic"
	"github.com/jhump/protoreflect/grpcreflect"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection/grpc_reflection_v1alpha"
)

type Proxy struct {
	cc         *grpc.ClientConn
	reflector  reflection.Reflector
	stub       stub.Stub
	descSource grpcurl.DescriptorSource
}

func main(){
	conn,err:=grpc.Dial(":50001",grpc.WithInsecure())
	if err!=nil{
		panic(err)
	}
	ctx:=context.TODO()
	rc:=grpcreflect.NewClient(ctx,grpc_reflection_v1alpha.NewServerReflectionClient(conn))
	proxy:=&Proxy{
		cc:         conn,
		reflector:  reflection.NewReflector(rc),
		stub:       stub.NewStub(grpcdynamic.NewStub(conn)),
		descSource: grpcurl.DescriptorSourceFromServer(ctx, rc),
	}
	md := make(metadata.Metadata)
	resp,err:=proxy.Invoke(ctx,"hello.SayHelloService","Hello",[]byte(`{"name":"dynamic","age":123}`),&md)
	if err!=nil{
		panic(err)
	}
	fmt.Print(string(resp))
}
func (p *Proxy) Invoke(ctx context.Context,
	serviceName, methodName string,
	message []byte,
	md *metadata.Metadata,
) ([]byte, error) {
	invocation, err := p.reflector.CreateInvocation(serviceName, methodName, message)
	if err != nil {
		return nil, err
	}

	outputMsg, err := p.stub.InvokeRPC(ctx, invocation, md)
	if err != nil {
		return nil, err
	}
	m, err := outputMsg.MarshalJSON()
	if err != nil {
		return nil, errors.New("failed to marshal output JSON")
	}
	return m, err
}