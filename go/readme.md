# 开发说明

调试命令

1.  启动调试

在当前目录下执行`fresh`, 会启动对当前目录 go 代码的自动监控, 变动后会自动重新编译

执行`./.tmp/app-go2fe`, 即可执行最新编译后的结果

# go2fe 使用指南

1.  安装项目`go get github.com/GO2F/Go2Fe`
2.  在 main.go 文件中引入`go2fe`包, 在`gin.Run`之前, 添加`go2fe.Config()`方法(方便注册所有接口)
    1.  在每个接口文件里, 通过`go2fe.RegistModel`注册数据模型, 或通过`DataModel`统一注册数据模型
    2.  添加`client`目录, 作为`go2fe`生成的前端代码所在地
    3.  添加`***`代码, 配置`static`为静态资源目录, 添加基础路由代码`***`, 默认返回其下的`index.html`作为入口文件
3.  执行`go run main.go go2fe:generate`, 生成前端代码
    1.  在 config 方法中, 监控命令行参数, 当发现命令行参数中包含`go2fe:generate`时, 启动页面创建流程
    2.  首先初始化项目文件夹, 在项目中创建`client`, `static`两个目录
    3.  然后生成 json 配置, 输出到`client/go2fe.json`目录中
    4.  然后切换到`client`目录, 下载`go2fe`的 npm 包, 执行 install
    5.  根据`go2fe.json`配置, 填充代码模板, 生成前端代码
    6.  执行`npm run build`, 构建前端项目, 将构建结果输出到`static`目录中
    7.  执行完毕
4.  运行程序

# todo

- [ ] 提供 crud 的接口返回格式规范
