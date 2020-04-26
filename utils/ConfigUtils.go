package utils

import (
	"fmt"
	"github.com/Unknwon/goconfig"
)

var cfg *goconfig.ConfigFile

func GetConfigIni(filepath string) (err error) {
	config, err := goconfig.LoadConfigFile(filepath)
	if err != nil {
		fmt.Println("配置文件读取错误,找不到配置文件", err)
		return err
	}
	cfg = config
	return nil
}
func init() {
	GetConfigIni("config/config.ini")
}

// 加载数据库配置
func GetDatabase() (types, local, prod string, maxOpenConns string, maxIdConns string, connMaxLifetime string, err error) {

	if types, err = cfg.GetValue("database", "types"); err != nil {
		fmt.Println("配置文件中不存在types", err)
		return types, local, prod, maxOpenConns, maxIdConns, connMaxLifetime, err
	}
	if local, err = cfg.GetValue("database", "dev"); err != nil {
		fmt.Println("配置文件中不存在dev", err)
		return types, local, prod, maxOpenConns, maxIdConns, connMaxLifetime, err
	}
	if prod, err = cfg.GetValue("database", "prod"); err != nil {
		fmt.Println("配置文件中不存在prod", err)
		return types, local, prod, maxOpenConns, maxIdConns, connMaxLifetime, err
	}

	if maxOpenConns, err = cfg.GetValue("database", "maxOpenConns"); err != nil {
		fmt.Println("配置文件中不存在prod", err)
		return types, local, prod, maxOpenConns, maxIdConns, connMaxLifetime, err
	}

	if maxIdConns, err = cfg.GetValue("database", "maxIdConns"); err != nil {
		fmt.Println("配置文件中不存在prod", err)
		return types, local, prod, maxOpenConns, maxIdConns, connMaxLifetime, err
	}
	if connMaxLifetime, err = cfg.GetValue("database", "connMaxLifetime"); err != nil {
		fmt.Println("配置文件中不存在prod", err)
		return types, local, prod, maxOpenConns, maxIdConns, connMaxLifetime, err
	}
	return types, local, prod, maxOpenConns, maxIdConns, connMaxLifetime, err
}

//加载服务器配置
func GetServer() (port string, err error) {
	if port, err = cfg.GetValue("server", "port"); err != nil {
		fmt.Println("配置文件中不存在port", err)
		return port, err
	}
	return port, nil
}

// 加载jwt配置
func GetJWT() (oneDayOfHours string, secret string, err error) {
	if oneDayOfHours, err = cfg.GetValue("JWT", "oneDayOfHours"); err != nil {
		fmt.Println("配置文件中不存在port", err)
		return oneDayOfHours, secret, err
	}
	if secret, err = cfg.GetValue("JWT", "secret"); err != nil {
		fmt.Println("配置文件中不存在port", err)
		return oneDayOfHours, secret, err
	}
	return oneDayOfHours, secret, err
}
