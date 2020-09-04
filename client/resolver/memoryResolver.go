package resolver

import (
	"fmt"
	"google.golang.org/grpc/resolver"
	"time"
)

type memoryResolver struct {
	target resolver.Target
	cc resolver.ClientConn
	addrsStore map[string][]string
}

func (m *memoryResolver) ResolveNow(resolver.ResolveNowOptions) {
	fmt.Println("ResolveNow")
	fmt.Println(time.Now())
}

func (m *memoryResolver) Close() {

}


var addrs=[]string{"localhost:50001","localhost:50001"}
func NewMemoryResolver(target resolver.Target,cc resolver.ClientConn) *memoryResolver{
	return &memoryResolver{target:target,cc:cc,
		addrsStore: map[string][]string{
			"demo-svc":addrs,
		},
		}
}
func (m *memoryResolver) start(){
	addrStore:=m.addrsStore[m.target.Endpoint]
	addrs:=make([]resolver.Address,len(addrStore))
	for i,v:=range addrStore{
		addrs[i]=resolver.Address{Addr:v}
	}
	m.cc.UpdateState(resolver.State{Addresses:addrs})
	go func() {
		time.Sleep(time.Second*10)
		m.cc.UpdateState(resolver.State{Addresses:[]resolver.Address{resolver.Address{Addr:addrStore[1]}}})
	}()
}