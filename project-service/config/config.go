package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
)

var AppConf = InitConfig()

type Config struct {
	viper        *viper.Viper
	ServerConfig *ServerConfig
	GrpcConfig   *GrpcConfig
	MqConfig     *MqConfig
	MysqlConfig  *MysqlConfig
}

type ServerConfig struct {
	Name string
	Addr string
}

type GrpcConfig struct {
	Name string
	Addr string
}

type MqConfig struct {
	Addrs []string
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
	workDir, _ := os.Getwd()
	fmt.Println(workDir)
	conf.viper.SetConfigName("config")
	conf.viper.SetConfigType("yaml")
	conf.viper.AddConfigPath(workDir + "/config")
	err := conf.viper.ReadInConfig()
	if err != nil {
		log.Fatalln(err)
	}
	conf.InitServerConfig()
	conf.InitEtcdConfig()
	return conf
}

func (c *Config) InitServerConfig() {
	sc := &ServerConfig{}
	sc.Name = c.viper.GetString("server.name")
	sc.Addr = c.viper.GetString("server.addr")
	c.ServerConfig = sc
}

func (c *Config) InitEtcdConfig() {
	mc := &MqConfig{}
	var addrs []string
	err := c.viper.UnmarshalKey("etcd.addrs", &addrs)
	if err != nil {
		log.Fatalln("init Etcd Config err: ", err)
	}
	mc.Addrs = addrs
	c.MqConfig = mc
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
