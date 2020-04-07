# 开发说明

# 项目启动

执行`go run main.go`, 自动安装依赖

# 测试项目

执行`go test`, 会自动自行`main_test.go`中以`Test`开头的函数, 可参考[这里](https://books.studygolang.com/The-Golang-Standard-Library-by-Example/chapter09/09.1.html)

# go2fe 使用指南

> 模板示例项目: [demo-project](https://github.com/GO2F/demo-project)

1.  安装依赖包`go get github.com/GO2F/Go2Fe`
2.  初始化项目

    1.  在 main.go 中引入

        1.  `go2fe "github.com/GO2F/Go2Fe"`
        2.  注册数据模型

            1.  通过`go2fe.RegModel()`注册数据模型
            2.  示例

                ```go
                // ComponentModel demo数据模型
                type ComponentModel struct {
                    ID          string `json:"ID" unique_key:"" show:"" title:"id"`
                    DisplayName string `json:"DisplayName" show:"" title:"组件名"`
                    PackageName string `json:"PackageName" show:"" title:"包名"`
                    DevListJSON string `json:"DevListJSON" show:"" title:"开发者"`
                    Description string `json:"Description" show:"" title:"描述"`
                    SiteURL     string `json:"SiteURL" show:"" title:"网站主页"`
                    Remark      string `json:"Remark" title:"备注"`
                }

                func main() {
                    customerModel := go2fe.ModelDefine{
                        DataModel: ComponentModel{},
                        Page: go2fe.Page{
                            Create: true,
                            Update: true,
                            Detail: true,
                        },
                        Name:        "测试模型",
                        BasePath:    "/component",
                        BaseAPIPath: "/api/component",
                    }
                    go2fe.RegModel(customerModel)

                    if go2fe.Bootstrap() {
                        return
                    }
                    dbErr := model.InitDb()
                    log.Init(config.App.LogPathURI)
                    // log.Info("进程以server模式启动")
                    if dbErr != nil {
                        // log.Error("数据库初始化失败,程序退出")
                        return
                    }
                    r := router.InitRouter()

                    // 程序结束前关闭数据库链接
                    defer model.DB.Close()
                    r.Run(":" + config.App.Port)
                }
                ```

            3.  数据模型字段说明
                1.  test

        3.  添加接口代码
            1.  每类数据模型需要实现以下五个接口, 参数定义见[这里](Swagger-openAPI描述)
                1.  get
                2.  list
                3.  create
                4.  update
                5.  remove
        4.  在 main 方法中添加
            1.  go2fe.Bootstrap()
        5.  执行 `go run main.go go2fe:init`, 初始化文件夹
            1.  或者直接使用内置函数启动`go2fe.InitFeTemplate()`
        6.  执行 `go run main.go go2fe:start`, 启动前端测试环境, 生成的前端代码位于`client`文件夹中
            1.  通过异步函数启动 dev 代码
            2.  或者直接使用内置函数启动`go2fe.StartDev()`
        7.  执行 `go run main.go go2fe:build`, 构建前端代码, 生成的静态资源位于`static`文件夹中
            1.  或者直接使用内置函数启动`go2fe.StartBuild()`

# 项目文件夹目录

- client
  - 前端代码目录
  - 需要添加到 .gitignore 中
- static
  - 静态资源目录
  - 需要添加到 .gitignore 中

# todo

- [ ] 提供 crud 的接口返回格式规范
