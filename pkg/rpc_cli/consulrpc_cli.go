package rpc_cli

import (
	"context"
	"fmt"
	"lightim/internal/discovery/resolver"
	"lightim/pkg/logger"
	"lightim/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"log"
)


var (
	LogicConSulIntClient   pb.LogicIntClient
	ConnectConSulIntClient pb.ConnIntClient
)

func InitConSulLogicIntClient(addr string,servicename string){

	schema, err := resolver.GenerateAndRegisterConsulResolver(addr, servicename)
	if err != nil {
		log.Fatal("init consul resovler err", err.Error())
	}
	conn, err := grpc.DialContext(context.TODO(),fmt.Sprintf("%s:///%s", schema,servicename), grpc.WithInsecure(),grpc.WithUnaryInterceptor(interceptor))

	if err != nil {
		logger.Sugar.Error(err)
		panic(err)
	}

	LogicIntClient = pb.NewLogicIntClient(conn)
}

func InitConSulConnIntClient(addr string,servicename string) {

	schema, err := resolver.GenerateAndRegisterConsulResolver(addr, servicename)
	if err != nil {
		log.Fatal("init consul resovler err", err.Error())
	}
	conn, err := grpc.DialContext(context.TODO(),fmt.Sprintf("%s:///%s", schema,servicename), grpc.WithInsecure(), grpc.WithBalancerName(roundrobin.Name),grpc.WithUnaryInterceptor(interceptor))

	if err != nil {
		logger.Sugar.Error(err)
		panic(err)
	}

	ConnectIntClient = pb.NewConnIntClient(conn)
}
