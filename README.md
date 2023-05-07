# general-service
1.基于 gin 的 web 服务
2.在顶层目录下执行 go get -d -v ./... 可以加载包的内容
go mode
1.go mod init 初始化模块
2.go get 添加依赖
杂记
1.os.GetWd()获取文件夹路径的方案

go 文件打包的指令
CGO_ENABLED=0  GOOS=linux  GOARCH=amd64  go build -o general-service main.go

docker 打包的指令
docker build -t general:v0.1 .  