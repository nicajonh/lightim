package main

import (
	"lightim/api/logic"
	"lightim/config"
	"lightim/internal/logic/db"
	"lightim/pkg/logger"
	"lightim/pkg/rpc_cli"
	"lightim/pkg/util"
)

func main() {
	// 初始化数据库
	db.InitDB()

	// 初始化自增id配置
	util.InitUID(db.DBCli)

	// 初始化RpcClient
	rpc_cli.InitConnIntClient(config.LogicConf.ConnRPCAddrs)

	/*// 启动nsq消费服务
	go func() {
		defer util.RecoverPanic()
		consume.StartNsqConsumer()
	}()
	*/

	logic.StartRpcServer()
	logger.Logger.Info("logic server start")
	select {}
}
