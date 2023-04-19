package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

var AppConf = InitConfig()

type Config struct {
	viper        *viper.Viper
	ServerConfig *ServerConfig
	EtcdConfig   *EtcdConfig
	GrpcConfig   *GrpcConfig
	MysqlConfig  *MysqlConfig
}

type ServerConfig struct {
	Name string
	Addr string
}

type EtcdConfig struct {
	Addrs []string
}

type GrpcConfig struct {
	Name    string
	Addr    string
	Version string
	Weight  int64
}

type MysqlConfig struct {
	Username string
	Password string
	Host     string
	Port     int
	Db       string
}

func InitConfig() *Config {
	conf := &Config{viper: viper.New()}
	workDir, _ := os.Getwd() //当前工作路径
	conf.viper.SetConfigName("config")
	conf.viper.SetConfigType("yaml")
	conf.viper.AddConfigPath(workDir + "/config")
	err := conf.viper.ReadInConfig()
	if err != nil {
		log.Fatalln(err)
	}

	conf.InitServerConfig()
	conf.InitEtcdConfig()
	conf.InitGrpcConfig()
	conf.InitMySqlConfig()

	return conf
}

func (c *Config) InitServerConfig() {
	sc := &ServerConfig{}
	sc.Name = c.viper.GetString("server.name")
	sc.Addr = c.viper.GetString("server.addr")
	c.ServerConfig = sc
}

func (c *Config) InitEtcdConfig() {
	ec := &EtcdConfig{}
	var addrs []string
	err := c.viper.UnmarshalKey("etcd.addrs", &addrs)
	if err != nil {
		log.Fatalln("init etcd failed: ", err)
	}
	ec.Addrs = addrs
	c.EtcdConfig = ec
}

func (c *Config) InitGrpcConfig() {
	gc := &GrpcConfig{}
	gc.Addr = c.viper.GetString("grpc.addr")
	gc.Name = c.viper.GetString("grpc.name")
	gc.Version = c.viper.GetString("grpc.version")
	gc.Weight = c.viper.GetInt64("grpc.weight")
	c.GrpcConfig = gc
}

func (c *Config) InitMySqlConfig() {
	ms := &MysqlConfig{}
	ms.Username = c.viper.GetString("mysql.username")
	ms.Password = c.viper.GetString("mysql.password")
	ms.Host = c.viper.GetString("mysql.host")
	ms.Port = c.viper.GetInt("mysql.port")
	ms.Db = c.viper.GetString("mysql.db")
	c.MysqlConfig = ms
}
