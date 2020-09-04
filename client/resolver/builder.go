package resolver

import "google.golang.org/grpc/resolver"

type resolverBuilder struct {

}

func init(){
	resolver.Register(&resolverBuilder{})
}

func (m *resolverBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	r:=NewMemoryResolver(target,cc)
	r.start()
	return r,nil
}

func (m *resolverBuilder) Scheme() string {
	return "demo"
}

func NewResolverBuilder() resolver.Builder{
	return &resolverBuilder{}
}