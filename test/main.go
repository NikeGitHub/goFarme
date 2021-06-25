package main

import (
	"code.shomes.cn/pkg/bootstrap/test"
	"code.shomes.cn/tfblue/smartgy/test/apis"
)

func main() {
	apis.NICKInit()
	test.UserKey = map[int32]string{
		1: "454e7818",
	}
	test.AppendToApi(apis.APINICK)
	test.Start()
}
