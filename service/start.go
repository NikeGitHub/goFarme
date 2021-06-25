package service

import (
	"fmt"

	bsGRpc "code.shomes.cn/pkg/bootstrap/grpc"
	"code.shomes.cn/tfblue/smartgy/protobuf"
	"code.shomes.cn/tfblue/smartgy/service/service"
	"google.golang.org/grpc"
)

func Start() {
	err := bsGRpc.Start(func(s *grpc.Server) {
		protobuf.RegisterSelectOptionServer(s, &service.SelectOption{})
	})

	fmt.Println("错误", err)
}
