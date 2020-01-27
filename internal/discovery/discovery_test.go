package discovery

import (
	"context"
	"lightim/internal/discovery/proto"
	"lightim/internal/discovery/register"
	"fmt"
	"lightim/internal/discovery/resolver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"testing"
	"time"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *proto.HelloRequest) (*proto.HelloResponse, error) {
	fmt.Println("client called! 8081")
	return &proto.HelloResponse{Result: "hi," + in.Name + "!"}, nil
}

const (
	host        = "127.0.0.1"
	port        = 8081
	consul_port = 8500
)
func TestDiscovery_Client(t *testing.T){
	schema, err := resolver.GenerateAndRegisterConsulResolver("127.0.0.1:8500", "HelloService")
	if err != nil {
		log.Fatal("init consul resovler err", err.Error())
	}

	// Set up a connection to the server.
	conn, err := grpc.Dial(fmt.Sprintf("%s:///HelloService", schema), grpc.WithInsecure(), grpc.WithBalancerName(roundrobin.Name))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := proto.NewHelloServiceClient(conn)

	// Contact the server and print out its response.
	name := "mysixtools"

	for {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		r, err := c.SayHello(ctx, &proto.HelloRequest{Name: name})
		if err != nil {
			log.Println("could not greet: %v", err)

		} else {
			log.Printf("Hello: %s", r.Result)
		}
		time.Sleep(time.Second)
	}

}
func TestDiscovery_Server(t *testing.T) {
	listen,err :=net.ListenTCP("tcp",&net.TCPAddr{net.ParseIP(host),port,""})
	if err!=nil{
		t.Error(err)
	}
	s:=grpc.NewServer()
	//register Server
	cr := register.NewConsulRegister(fmt.Sprintf("%s:%d", host, consul_port), 15)
	cr.Register(&RegisterInfo{
		Host:           host,
		Port:           port,
		ServiceName:    "HelloService",
		UpdateInterval: time.Second})

	proto.RegisterHelloServiceServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(listen); err != nil {
		fmt.Println("failed to serve:" + err.Error())
	}
}
