package db

import (
	"database/sql"
	"fmt"
	"lightim/config"
	_ "github.com/go-sql-driver/mysql"
)

var DBCli *sql.DB

func init() {
	var err error
	mysqlConfig:=config.Conf.Logic.LogicMysql
	connArgs := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",  mysqlConfig.MysqlUser,mysqlConfig.MysqlPassword,mysqlConfig.MysqlAddress, mysqlConfig.MysqlPort, mysqlConfig.MysqlDb)
	DBCli, err = sql.Open("mysql",connArgs)
	if err != nil {
		panic(err)
	}
}
