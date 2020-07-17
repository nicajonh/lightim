package logic

import (
	"lightim/config"
	"lightim/pkg/pb"
	"lightim/pkg/util"
	"google.golang.org/grpc"
	"net"
)

// StartRpcServer 启动rpc服务
func StartRpcServer() {
	go func() {
		defer util.RecoverPanic()

		intListen, err := net.Listen("tcp", config.Conf.Logic.LogicRpcConf.RprClientExtListenAddr)
		if err != nil {
			panic(err)
		}
		intServer := grpc.NewServer(grpc.UnaryInterceptor(LogicClientExtInterceptor))
		pb.RegisterLogicForClientExtServer(intServer,&LogicClientExtServer{})
		err = intServer.Serve(intListen)
		if err != nil {
			panic(err)
		}
	}()

	go func() {
		defer util.RecoverPanic()

		extListen, err := net.Listen("tcp", config.Conf.Logic.LogicRpcConf.RpcConnListenAddr)
		if err != nil {
			panic(err)
		}
		extServer := grpc.NewServer(grpc.UnaryInterceptor(LogicIntInterceptor))
		pb.RegisterLogicForConnExtServer(extServer, &LogicForConnServer{})
		err = extServer.Serve(extListen)
		if err != nil {
			panic(err)
		}
	}()

	go func() {
		defer util.RecoverPanic()

		intListen, err := net.Listen("tcp", config.Conf.Logic.LogicRpcConf.RpcServerExtListenAddr)
		if err != nil {
			panic(err)
		}
		intServer := grpc.NewServer(grpc.UnaryInterceptor(LogicServerExtInterceptor))
		pb.RegisterLogicForServerExtServer(intServer, &LogicServerExtServer{})
		err = intServer.Serve(intListen)
		if err != nil {
			panic(err)
		}
	}()
}
//
//func example(){
//	listen,err :=net.ListenTCP("tcp",&net.TCPAddr{net.ParseIP(host),port,""})
//	if err!=nil{
//		t.Error(err)
//	}
//	s:=grpc.NewServer()
//	//register Server
//	cr := register.NewConsulRegister(fmt.Sprintf("%s:%d", host, consul_port), 15)
//	cr.Register(&RegisterInfo{
//		Host:           host,
//		Port:           port,
//		ServiceName:    "HelloService",
//		UpdateInterval: time.Second})
//
//	proto.RegisterHelloServiceServer(s, &server{})
//	reflection.Register(s)
//	if err := s.Serve(listen); err != nil {
//		fmt.Println("failed to serve:" + err.Error())
//	}
//	}