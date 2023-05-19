package common

import (
	"fmt"
	"gopkg.in/ini.v1"
	"strconv"
)

func init() {
	initConfig()
}

var (
	DataSource    string
	RedisStr      string
	JwtSecret     string
	JwtSecretByte []byte
)

func initConfig() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		fmt.Printf("无法加载配置文件：%v\n", err)
		return
	}

	// 读取指定 section 的指定 key 的值
	DataSource = cfg.Section("system").Key("dataSource").String()
	RedisStr = cfg.Section("system").Key("redisStr").String()

	// ===== 创建秘钥
	JwtSecret = cfg.Section("system").Key("jwtSecret").String()
	JwtSecretByte = []byte(JwtSecret)

	_, _ = strconv.Atoi(cfg.Section("system").Key("port").String())

	fmt.Println("aa")
}
