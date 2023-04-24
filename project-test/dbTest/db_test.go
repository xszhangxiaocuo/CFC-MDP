package dbTest

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

//var dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", "m_sec", "m_sec13674861950!", "rm-2vccslz6ay962myxuxo.mysql.cn-chengdu.rds.aliyuncs.com", 3307, "m_sec")

func Test1(t *testing.T) {
	username := "xszxc"                                       //账号
	password := "Qaz6659644."                                 //密码
	host := "gz-cynosdbmysql-grp-diyqqtkv.sql.tencentcdb.com" //数据库地址，可以是Ip或者域名
	port := "23946"                                           //数据库端口
	Dbname := "mdp_user"                                      //数据库名
	//maxOpenConns := config.Conf.MysqlConfig.MaxOpenConns //连接池最大连接数
	//maxIdleConns := config.Conf.MysqlConfig.MaxIdleConns // 连接池最大空闲连接数
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, Dbname)
	// 连接数据库
	_, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(fmt.Errorf("cannot establish db connection: %w", err))
	} else {
		fmt.Println("success")
	}

}
