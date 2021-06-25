package main

import (
	"fmt"

	"code.shomes.cn/pkg/bootstrap"
	"code.shomes.cn/pkg/bootstrap/tool"

	"code.shomes.cn/pkg/log"
	"code.shomes.cn/tfblue/smartgy/api"
	"code.shomes.cn/tfblue/smartgy/service"
)

func main() {
	log.New(&log.Options{
		File: &log.FileCfg{
			Filename:   "tfl.log",
			MaxSize:    512,
			MaxBackups: 100,
			MaxAge:     30,
			Compress:   true,
		},
	})
	b := bootstrap.NewBoot(&tool.BootOption{
		ProgramName: "smartgy",
		DB:          []string{"smartgy"},
		GRpcTimeout: 60,
	})
	b.Register("service", service.Start)
	b.Register("api", api.Start)
	err := b.Start()
	fmt.Println(err)
	fmt.Println("可选服务: api,service,cron")
}
