# service

服务

- 启动: go run main.go -c service
- 测试: go run test/main.go -a {api}/xxx
- 编译: protobuf: protoc -I ./protobuf/ --go_out=plugins=grpc,paths=source_relative:./protobuf/ ./protobuf/*.proto
- 自动重载: air
