package main

import (
	"Ditributed-Crontab/master"
	"flag"
	"fmt"
)

var (
	confFile string // 配置文件路径
)

// 解析命令行参数
func initArgs() {
	// master -config ./master.json -xxx 123 -yyy ddd
	// master -h
	flag.StringVar(&confFile, "config", "./master.json", "指定master.json")
	flag.Parse()
}

func main() {
	var (
		err error
	)
	//加载配置
	if err = master.InitConfig(confFile); err != nil {
		goto ERR
	}
	//启动任务管理器
	if err = master.InitJobMgr(); err != nil {
		goto ERR
	}
	//启动API HTTP服务
	if err = master.InitApiServer(); err != nil {
		goto ERR
	}
ERR:
	fmt.Println(err)
}
