package tcp_conn

import (
	"context"
	"lightim/config"
	"lightim/internal/tcp_conn"
	"lightim/pkg/logger"
	"lightim/pkg/pb"
	"net"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type ConnForLogicServer struct{}

// Message 投递消息
func (s *ConnForLogicServer) DeliverMessage(ctx context.Context, req *pb.DeliverMessageReq) (*pb.DeliverMessageResp, error) {
	return &pb.DeliverMessageResp{}, tcp_conn.DeliverMessage(ctx, req)
}

// UnaryServerInterceptor 服务器端的单向调用的拦截器
func UnaryServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	resp, err := handler(ctx, req)
	logger.Logger.Debug("interceptor", zap.Any("info", info), zap.Any("req", req), zap.Any("resp", resp))
	return resp, err
}

// StartRPCServer 启动rpc服务器
func StartRPCServer() {
	tcpConfig:=config.Conf.Connect.ConnTcpConf
	listener, err := net.Listen("tcp", tcpConfig.RPCListenAddr)
	if err != nil {
		panic(err)
	}
	server := grpc.NewServer(grpc.UnaryInterceptor(UnaryServerInterceptor))
	pb.RegisterConnForLogicExtServer(server, &ConnForLogicServer{})
	logger.Logger.Debug("rpc服务已经开启")
	err = server.Serve(listener)
	if err != nil {
		panic(err)
	}
}
