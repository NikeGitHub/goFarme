package apis

import (
	"code.shomes.cn/pkg/bootstrap/test"
	"code.shomes.cn/tfblue/smartgy/protobuf"
)

func NICKInit() {
	test.Env = "local"
	test.EnvApi = map[string]string{
		"local": "http://127.0.0.1:8347",
		"test":  "http://106.75.154.221:8359",
		"pro":   "http://106.75.154.221:8345",
	}
}

var (
	APINICK = map[string][]test.Param{
		"{api}/smart/select-option/project-use": {{UserId: 1, Param: protobuf.SelectOptionProjectUseIn{}, Comment: "工程用途下拉列表"}},
	}
)
