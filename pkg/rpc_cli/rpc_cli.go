package rpc_cli

import (
	"context"
	"fmt"
	"lightim/pkg/grpclib"
	"lightim/pkg/logger"
	"lightim/pkg/pb"
	"google.golang.org/grpc"
)

var (
	LogicForConnExtClient   pb.LogicForConnExtClient
	ConnForLogicExtClient pb.ConnForLogicExtClient
)

func InitLogicIntClient(addr string) {
	conn, err := grpc.DialContext(context.TODO(), addr, grpc.WithInsecure(), grpc.WithUnaryInterceptor(interceptor))
	if err != nil {
		logger.Sugar.Error(err)
		panic(err)
	}
	LogicForConnExtClient = pb.NewLogicForConnExtClient(conn)
}

func InitConnIntClient(addr string) {
	conn, err := grpc.DialContext(context.TODO(), addr, grpc.WithInsecure(), grpc.WithUnaryInterceptor(interceptor),
		grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, grpclib.Name)))
	if err != nil {
		logger.Sugar.Error(err)
		panic(err)
	}

	ConnForLogicExtClient = pb.NewConnForLogicExtClient(conn)
}
