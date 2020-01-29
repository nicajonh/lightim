package main

import (
	"lightim/api/tcp_conn"
	"lightim/config"
	tcp_conn2 "lightim/internal/tcp_conn"
	"lightim/pkg/rpc_cli"
	"lightim/pkg/util"
)

func main() {
	// 启动rpc服务
	go func() {
		defer util.RecoverPanic()
		tcp_conn.StartRPCServer()
	}()
	connConf:=config.Conf.Connect
	// 初始化Rpc Client
	rpc_cli.InitLogicIntClient(connConf.ConnTcpConf.LogicRPCAddrs)

	// 启动长链接服务器
	server := tcp_conn2.NewTCPServer(connConf.ConnTcpConf.TCPListenAddr, 10)
	server.Start()
}
