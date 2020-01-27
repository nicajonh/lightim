package config

import (
	"lightim/pkg/logger"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"os"
	"path/filepath"
	"sync"
)

var once sync.Once
var realpath string
var Conf *Config

type Config struct {
	Connect ConnConfig
	Logic   LogicConfig
	Site    SiteConfig
}

func init() {
	Init()
}


func Init() {
	once.Do(func() {
		env := GetMode()
		SetLogLevel(env)
		realPath, _ := filepath.Abs("./")
		configFilePath := realPath + "/config/" + env + "/"
		viper.SetConfigType("toml")
		viper.SetConfigName("/connect")
		viper.AddConfigPath(configFilePath)
		err := viper.ReadInConfig()
		if err != nil {
			panic(err)
		}
		viper.SetConfigName("/logic")
		err = viper.MergeInConfig()
		if err != nil {
			panic(err)
		}
		viper.SetConfigName("/site")
		err = viper.MergeInConfig()
		if err != nil {
			panic(err)
		}
		Conf = new(Config)
		viper.Unmarshal(&Conf.Connect)
		viper.Unmarshal(&Conf.Logic)
		viper.Unmarshal(&Conf.Site)
	})
}


func GetMode() string{
	env:=os.Getenv("RUN_MODE")
	if env==""{
		env="dev"
	}
	return env
}

func SetLogLevel(env string){
	switch env {
		case "dev":
			logger.Leavel = zap.DebugLevel
			logger.Target = logger.Console
		case "test":
			logger.Leavel = zap.DebugLevel
			logger.Target = logger.File
		case "prod":
			logger.Leavel = zap.InfoLevel
			logger.Target = logger.File
		default:
			logger.Leavel = zap.DebugLevel
			logger.Target = logger.Console
	}
}


// logic配置
type LogicConfig struct {
	LogicConsul LogicConsul `mapstructure:"logic-consul"`
	LogicRedis LogicRedis `mapstructure:"logic-redis"`
	LogicNsq LogicNsq `mapstructure:"logic-nsq"`
	LogicMysql LogicMysql  `mapstructure:"logic-mysql"`
	LogicRpcConf LogicRpcConf `mapstructure:"logic-rpc"`
}

type LogicRpcConf struct{
	RpcConnListenAddr       string `mapstructure:"rpcConnListenAddr"`
	RprClientExtListenAddr  string `mapstructure:"rprClientExtListenAddr"`
	RpcServerExtListenAddr  string `mapstructure:"rpcServerExtListenAddr"`
	ConnRPCAddrs            string `mapstructure:"connRPCAddrs"`
}

type LogicNsq struct{
	NsqHost string  `mapstructure:"nsqHost"`
	NsqPort int     `mapstructure:"nsqPort"`
}

type LogicConsul struct {
	ConsulHost             string `mapstructure:"consulHost"`
	ConsulPort             int    `mapstructure:"consulPort"`
}
type LogicRedis struct{
	RedisAddress  string `mapstructure:"redisAddress"`
	RedisPassword string `mapstructure:"redisPassword"`
}
type LogicMysql struct{
	MysqlAddress string `mapstructure:"mysqlAddress"`
	MysqlUser string `mapstructure:"mysqlUser"`
	MysqlPassword string `mapstructure:"mysqlPassword"`
	MysqlPort int `mapstructure:"mysqlPort"`
	MysqlDb string `mapstructure:"mysqlDb"`

}

type ConnConfig struct {
	ConnTcpConf ConnTcpConf `mapstructure:"conn-tcp"`
	ConnWebsocket ConnWebsocket `mapstructure:"conn-websocket"`
}
// conn配置
type ConnTcpConf struct {
	TCPListenAddr string `mapstructure:"tcpListenAddr"`
	RPCListenAddr string `mapstructure:"rpcListenAddr"`
	LocalAddr     string `mapstructure:"localAddr"`
	LogicRPCAddrs string `mapstructure:"logicRpcAddrs"`
}

// WS配置
type ConnWebsocket struct {
	WSListenAddr  string `mapstructure:"wsListenAddr"`
	RPCListenAddr string `mapstructure:"rpcListenAddr"`
	LocalAddr     string `mapstructure:"localAddr"`
	LogicRPCAddrs string `mapstructure:"logicRpcAddrs"`
}

type SiteBase struct {
	ListenPort int `mapstructure:"listenPort"`
}

type SiteConfig struct {
	SiteBase SiteBase `mapstructure:"site-base"`
}



