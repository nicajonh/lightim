package logic

import (
	"fmt"
	"lightim/config"
	"lightim/internal/discovery"
	"lightim/internal/discovery/register"
	"lightim/pkg/pb"
	"lightim/pkg/util"
	"google.golang.org/grpc"
	"net"
	"time"
)

// StartConsulRpcServer 启动rpc服务
func StartConsulRpcServer() {
	go func() {
		defer util.RecoverPanic()
		logicConf:=config.Conf.Logic
		intListen, err := net.Listen("tcp", logicConf.LogicRpcConf.RpcConnListenAddr)
		if err != nil {
			panic(err)
		}
		intServer := grpc.NewServer(grpc.UnaryInterceptor(LogicIntInterceptor))

		cr := register.NewConsulRegister(fmt.Sprintf("%s:%d", logicConf.LogicConsul.ConsulHost, logicConf.LogicConsul.ConsulPort), 15)
		err=cr.Register(&discovery.RegisterInfo{
			Host:           logicConf.LogicConsul.ConsulHost,
			Port:           logicConf.LogicConsul.ConsulPort,
			ServiceName:    "LogicServer",
			UpdateInterval: time.Second,
		})
		if err != nil {
			panic(err)
		}
		pb.RegisterLogicForConnExtServer(intServer, &LogicForConnServer{})
		err = intServer.Serve(intListen)
		if err != nil {
			panic(err)
		}
	}()

	go func() {
		defer util.RecoverPanic()
		logicConf:=config.Conf.Logic
		extListen, err := net.Listen("tcp", logicConf.LogicRpcConf.RprClientExtListenAddr)
		if err != nil {
			panic(err)
		}
		extServer := grpc.NewServer(grpc.UnaryInterceptor(LogicClientExtInterceptor))
		cr := register.NewConsulRegister(fmt.Sprintf("%s:%d", logicConf.LogicConsul.ConsulHost, logicConf.LogicConsul.ConsulPort), 15)
		err=cr.Register(&discovery.RegisterInfo{
			Host:           logicConf.LogicConsul.ConsulHost,
			Port:           logicConf.LogicConsul.ConsulPort,
			ServiceName:    "LogicClientExt",
			UpdateInterval: time.Second,
		})
		if err != nil {
			panic(err)
		}
		pb.RegisterLogicForClientExtServer(extServer, &LogicClientExtServer{})
		err = extServer.Serve(extListen)
		if err != nil {
			panic(err)
		}
	}()

	go func() {
		defer util.RecoverPanic()
		logicConf:=config.Conf.Logic
		intListen, err := net.Listen("tcp", logicConf.LogicRpcConf.RpcServerExtListenAddr)
		if err != nil {
			panic(err)
		}
		intServer := grpc.NewServer(grpc.UnaryInterceptor(LogicServerExtInterceptor))
		cr := register.NewConsulRegister(fmt.Sprintf("%s:%d", logicConf.LogicConsul.ConsulHost, logicConf.LogicConsul.ConsulPort), 15)
		err=cr.Register(&discovery.RegisterInfo{
			Host:           logicConf.LogicConsul.ConsulHost,
			Port:           logicConf.LogicConsul.ConsulPort,
			ServiceName:    "LogicServerExt",
			UpdateInterval: time.Second,
		})
		if err != nil {
			panic(err)
		}
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