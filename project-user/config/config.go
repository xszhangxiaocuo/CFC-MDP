package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

var AppConfig = InitConfig()

type Config struct {
	viper        *viper.Viper
	ServerConfig *ServerConfig
	EtcdConfig   *EtcdConfig
}

type ServerConfig struct {
	Name string
	Addr string
}

type EtcdConfig struct {
	Addrs []string
}

func InitConfig() *Config {
	conf := &Config{viper: viper.New()}
	workDir, _ := os.Getwd() //当前工作路径
	conf.viper.SetConfigName("config")
	conf.viper.SetConfigType("ymal")
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
	ec := &EtcdConfig{}
	var addrs []string
	err := c.viper.UnmarshalKey("etcd.addrs", addrs)
	if err != nil {
		log.Fatalln("init etcd failed: ", err)
	}
	ec.Addrs = addrs
	c.EtcdConfig = ec
}
