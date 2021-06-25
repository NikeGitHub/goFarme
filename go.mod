module code.shomes.cn/tfblue/smartgy

go 1.15

require (
	code.shomes.cn/pkg/bootstrap v0.0.0-20210422111552-90cdb7b8cf3d
	code.shomes.cn/pkg/db v0.0.0-20210602102515-9f1d91161413
	code.shomes.cn/pkg/log v1.0.1
	github.com/gin-gonic/gin v1.7.2
	github.com/go-sql-driver/mysql v1.6.0
	github.com/golang/protobuf v1.5.2
	golang.org/x/net v0.0.0-20210614182718-04defd469f4e
	google.golang.org/grpc v1.38.0
	google.golang.org/grpc/examples v0.0.0-20210625051147-9b2fa9f8d3ca // indirect
)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
