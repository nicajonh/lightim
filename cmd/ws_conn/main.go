package main

import (
	"lightim/api/ws_conn"
	"lightim/config"
	ws_conn2 "lightim/internal/ws_conn"
	"lightim/pkg/rpc_cli"
	"lightim/pkg/util"
)

func main() {
	// 启动rpc服务
	go func() {
		defer util.RecoverPanic()
		ws_conn.StartRPCServer()
	}()
	connConf:=config.Conf.Connect
	// 初始化Rpc Client
	rpc_cli.InitLogicIntClient(connConf.ConnWebsocket.LogicRPCAddrs)

	// 启动长链接服务器
	ws_conn2.StartWSServer(connConf.ConnWebsocket.WSListenAddr)
}
