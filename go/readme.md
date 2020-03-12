# 开发说明

调试命令

1.  启动调试

在当前目录下执行`fresh`, 会启动对当前目录 go 代码的自动监控, 变动后会自动重新编译

执行`./.tmp/app-go2fe`, 即可执行最新编译后的结果

# go2fe 使用指南

1.  安装项目`go get github.com/GO2F/Go2Fe`
2.  在 main.go 文件中引入`go2fe`包, 在第一行添加`go2fe.Config()`方法
3.  执行`go run main.go go2fe:init`, 启动交互控制台
4.  执行`go run main.go go2fe:generate`, 直接生成前端代码
5.  运行程序
