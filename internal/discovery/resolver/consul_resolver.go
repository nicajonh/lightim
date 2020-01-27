package resolver

import (
	"context"
	"fmt"
	"lightim/internal/discovery"
	consulapi "github.com/hashicorp/consul/api"
	"google.golang.org/grpc/naming"
	"google.golang.org/grpc/resolver"
	"log"
	"sync"
	"time"
)

type consulBuilder struct{
	address string
	client *consulapi.Client
	serviceName string
}

func (self *consulBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOption) (resolver.Resolver, error) {
		self.serviceName=target.Endpoint
		//get service info and config
		addrs,serviceConfig,err :=self.resolve()
		if err!=nil{
			return nil,err
		}
		//update service info to client
		cc.NewAddress(addrs)
		cc.NewServiceConfig(serviceConfig)
		consulResolver := NewConsulResolver(&cc, self, opts)
		consulResolver.wg.Add(1)
		go consulResolver.watcher()
		return consulResolver,nil
}

func NewConsulBuilder(address string) resolver.Builder{
	client,err:=discovery.GenerateConsulClient(address)
	if err!=nil{
		log.Println("Create consul client error:", err.Error())
		return nil
	}
	return &consulBuilder{address:address,client:client}
}

//resovler
func(self *consulBuilder) resolve()([]resolver.Address,string,error){
	serviceEntries, _, err := self.client.Health().Service(self.serviceName, "", true, &consulapi.QueryOptions{})
	if err!=nil{
		return nil,"",err
	}
	addrs :=make([]resolver.Address,0)
	for _,serviceEntry:=range serviceEntries{
		address := resolver.Address{Addr: fmt.Sprintf("%s:%d", serviceEntry.Service.Address, serviceEntry.Service.Port)}
		addrs = append(addrs,address)
	}
	return addrs,"",nil
}

func (self *consulBuilder) Scheme() string {
	return "consul"
}

type consulResolver struct{
	clientConn *resolver.ClientConn
	consulBuilder *consulBuilder
	t *time.Ticker
	wg sync.WaitGroup
	rn chan struct{}
	ctx context.Context
	cancel context.CancelFunc
	disableServiceConfig bool
}

func NewConsulResolver(cc *resolver.ClientConn, cb *consulBuilder, opts resolver.BuildOption) *consulResolver {
	ctx, cancel := context.WithCancel(context.Background())
	return &consulResolver{
		clientConn:           cc,
		consulBuilder:        cb,
		t:                    time.NewTicker(time.Second),
		ctx:                  ctx,
		cancel:               cancel,
		disableServiceConfig: opts.DisableServiceConfig}
}

func(self *consulResolver) watcher(){
	self.wg.Done()
	for{
		select {
		case <-self.ctx.Done():
			return
		case <-self.rn:
		case <-self.t.C:
		}
		addrs,serviceConfig,err:=self.consulBuilder.resolve()
		if err!=nil{
			log.Fatal("query service entries error:", err.Error())
		}
		(*self.clientConn).NewServiceConfig(serviceConfig)
		(*self.clientConn).NewAddress(addrs)
	}
}


func (self *consulResolver) Scheme() string {
	return self.consulBuilder.Scheme()
}

func (self *consulResolver) ResolveNow(rno resolver.ResolveNowOption) {
	select {
	case self.rn <- struct{}{}:
	default:
	}
}

func (self *consulResolver) Close() {
	self.cancel()
	self.wg.Wait()
	self.t.Stop()
}

type consulClientConn struct {
	adds []resolver.Address
	sc   string
}



func NewConsulClientConn() resolver.ClientConn{
	return &consulClientConn{}
}

func (self *consulClientConn) NewAddress(addresses []resolver.Address) {
	self.adds = addresses
}

func (self *consulClientConn) NewServiceConfig(serviceConfig string) {
	self.sc = serviceConfig
}

func GenerateAndRegisterConsulResolver(address string,serviceName string)(scheme string,err error){
	builder := NewConsulBuilder(address)
	target := resolver.Target{Scheme: builder.Scheme(), Endpoint: serviceName}
	_, err = builder.Build(target, NewConsulClientConn(), resolver.BuildOption{})
	if err != nil {
		return builder.Scheme(), err
	}
	resolver.Register(builder)
	scheme=builder.Scheme()
	return
}